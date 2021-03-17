package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var mu sync.Mutex

const (
	iface   = "any"
	snapLen = int32(1500)
	promisc = false
	timeout = pcap.BlockForever
)

// Information about the packet
type data struct {
	firstSeen   string
	udpOrTcp    string
	srcIP       string
	srcPort     string
	dstIP       string
	dstPort     string
	totalAmount int
}

// Start prometheus listener.
func startPrometheus(port string) {
	n, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("error: failed to open prometheus listen port: %v\n", err)
		os.Exit(1)
	}
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.Handler())
	http.Serve(n, m)
}

// doMetrics will register all the metrics for IPMap
func doMetrics(IPMap map[string]map[string]data, refresh int) {
	hosts := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "hosts_src_dst",
			Help: "Number of bytes transfered between hosts",
		},
		// []string{"addr", "port", "firstSeen", "srcPort", "dstPort"},
		[]string{"addr", "port", "firstSeen", "dstPort"},
	)

	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(hosts)

	for {
		mu.Lock()
		for k1, v1 := range IPMap {
			for k2, v2 := range v1 {
				// hosts.With(prometheus.Labels{"addr": k1, "port": k2, "firstSeen": v2.firstSeen.String(), "srcPort": v2.srcPort, "dstPort": v2.dstPort}).Set(float64(v2.totalAmount))
				hosts.With(prometheus.Labels{"addr": k1, "port": k2, "firstSeen": v2.firstSeen, "dstPort": v2.dstPort}).Set(float64(v2.totalAmount))
			}
		}
		mu.Unlock()

		time.Sleep(time.Second * time.Duration(refresh))
	}
}

func main() {
	iface := flag.String("iface", "", "the name of the interface to listen on")
	filter := flag.String("filter", "", "filter to use, same as nmap filters")
	promHTTP := flag.String("promHTTP", ":8888", "set ip and port for prometheus to listen. Ex. localhost:8888")
	promRefresh := flag.Int("promRefresh", 5, "the refresh rate in seconds that prometheus should refresh the metrics")
	var localIPs flagStringSlice
	flag.Var(&localIPs, "localIPs", "comma separated list of local host adresses")
	flag.Parse()

	if *iface == "" {
		log.Printf("error: you have to specify an interface to listen on\n")
		os.Exit(1)
	}

	if !localIPs.ok {
		log.Printf("error: no local host ip's specified\n")
		os.Exit(1)
	}

	localIPMap := map[string]struct{}{}
	for _, v := range localIPs.values {
		localIPMap[v] = struct{}{}
	}

	go startPrometheus(*promHTTP)

	// Get a BPF filter handle that we can set the filter on.
	handle, err := pcap.OpenLive(*iface, 65535, true, pcap.BlockForever)
	if err != nil {
		log.Printf("error: pcap.OpenLive failed: %v\n", err)
	}
	defer handle.Close()

	err = handle.SetBPFFilter(*filter)
	if err != nil {
		log.Printf("error: handle.SetBPFFilter failed: %v\n", err)
	}

	IPMap := map[string]map[string]data{}

	go doMetrics(IPMap, *promRefresh)

	var eth layers.Ethernet
	var ip4 layers.IPv4
	var tcp layers.TCP
	var udp layers.UDP
	var payload gopacket.Payload
	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip4, &tcp, &udp, &payload)
	decoded := make([]gopacket.LayerType, 0, 10)

	for {
		packetData, ci, err := handle.ZeroCopyReadPacketData()
		if err != nil {
			log.Printf("error getting packet: %v %v", err, ci)
			continue
		}
		err = parser.DecodeLayers(packetData, &decoded)
		if err != nil {
			// log.Printf("error decoding packet: %v", err)
			continue
		}

		if err := parser.DecodeLayers(packetData, &decoded); err != nil {
			log.Printf("error: could not decode layers: %v\n", err)
		}

		d := data{}

		for _, typ := range decoded {
			switch typ {
			case layers.LayerTypeEthernet:
			case layers.LayerTypeIPv4:
				d.firstSeen = time.Now().Format("2006 01 2 15:04:05")
				d.srcIP = ip4.SrcIP.String()
				d.dstIP = ip4.DstIP.String()
			case layers.LayerTypeTCP:
				d.udpOrTcp = "tcp"
				d.srcPort = d.udpOrTcp + "/" + tcp.SrcPort.String()
				d.dstPort = d.udpOrTcp + "/" + tcp.DstPort.String()
			case layers.LayerTypeUDP:
				d.udpOrTcp = "udp"
				d.srcPort = d.udpOrTcp + "/" + udp.SrcPort.String()
				d.dstPort = d.udpOrTcp + "/" + udp.DstPort.String()
			case gopacket.LayerTypePayload:
				d.totalAmount = len(payload.LayerContents())
			}
		}

		if d.totalAmount == 0 {
			continue
		}

		key1srcDst := d.srcIP + "->" + d.dstIP + ", proto: " + d.udpOrTcp

		key1srcDstRev := d.dstIP + "->" + d.srcIP + ", proto: " + d.udpOrTcp
		// Check if this is the return traffic for udp
		if _, ok := IPMap[key1srcDstRev][d.srcPort]; ok {

			// Check if the ip where defined as a local ip at startup
			_, ok2 := localIPMap[d.dstIP]

			if ok2 || d.dstIP == "127.0.0.1" {
				d.dstPort = "reply_" + d.srcPort
			}
		}

		// If already present, copy totalLength and time from previous.
		if v, ok := IPMap[key1srcDst][d.dstPort]; ok {
			d.totalAmount = v.totalAmount + d.totalAmount
			d.firstSeen = v.firstSeen
		} else if v, ok := IPMap[key1srcDst]["reply_"+d.dstPort]; ok {
			d.totalAmount = v.totalAmount + d.totalAmount
			d.firstSeen = v.firstSeen
		}

		// Declare the inner port map, and then store it in the outer hosts map.
		protoMap := map[string]data{}
		protoMap[d.dstPort] = d
		mu.Lock()
		IPMap[key1srcDst] = protoMap
		mu.Unlock()

	}
}
