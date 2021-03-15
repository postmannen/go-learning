package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	iface   = "any"
	snapLen = int32(1600)
	promisc = false
	timeout = pcap.BlockForever
)

type data struct {
	firstSeen   time.Time
	udpOrTcp    string
	srcIP       string
	srcPort     string
	dstIP       string
	dstPort     string
	totalAmount int
}

func createMapValue(ipLayer gopacket.Layer, packet gopacket.Packet, IPMap map[string]map[string]data) {
	ip, _ := ipLayer.(*layers.IPv4)
	appLayer := packet.ApplicationLayer()
	if appLayer != nil {

		d := data{firstSeen: time.Now()}

		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			d.srcPort = tcp.SrcPort.String()
			d.dstPort = tcp.DstPort.String()
			d.udpOrTcp = "tcp"

			fmt.Printf("flow: #%v\n", tcp.TransportFlow().String())
		}

		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
			udp, _ := udpLayer.(*layers.UDP)
			d.srcPort = udp.SrcPort.String()
			d.dstPort = udp.DstPort.String()
			d.udpOrTcp = "udp"
		}

		//fmt.Printf("packet : %#v\n", packet)

		d.srcIP = ip.SrcIP.String()
		d.dstIP = ip.DstIP.String()
		key1srcDst := d.srcIP + "->" + d.dstIP + ", proto: " + d.udpOrTcp
		key2portInfo := d.udpOrTcp + "/" + d.dstPort

		// layerType := packet.ApplicationLayer().LayerType().String()
		d.totalAmount = packet.Metadata().Length

		// If already present, copy totalLength and time from previous.
		if v, ok := IPMap[key1srcDst]; ok && IPMap[key1srcDst][key2portInfo].udpOrTcp == d.udpOrTcp {
			d.totalAmount = v[key2portInfo].totalAmount + d.totalAmount
			d.firstSeen = v[key2portInfo].firstSeen
		}

		// Declare the inner map, and then store it in the outer map.
		protoMap := map[string]data{}
		protoMap[key2portInfo] = d
		IPMap[key1srcDst] = protoMap
	}
}

func printMap(IPMap map[string]map[string]data, timeStart time.Time) {
	fmt.Printf("--------------------Start: %v-----------------------\n", timeStart)

	for k, v := range IPMap {
		fmt.Printf("addr: %v", k)
		for k, v := range v {
			fmt.Printf(", port: %v", k)
			fmt.Printf(", size: %v, firstSeen: %v, srcPort: %v, dstPort: %v", v.totalAmount, v.firstSeen, v.srcPort, v.dstPort)
		}
		fmt.Println()
	}
	fmt.Printf("--------------------------------------------\n")
}

func startPrometheus() {
	n, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Printf("error: failed to open prometheus listen port: %v\n", err)
		os.Exit(1)
	}
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.Handler())
	http.Serve(n, m)
}

func main() {

	go startPrometheus()

	filter := flag.String("filter", "", "filter to use, same as nmap filters")
	flag.Parse()

	// Check if the interface exists, or is set to "any"
	ifs, err := pcap.FindAllDevs()
	if err != nil {
		log.Println("error: pcap.FindAllDevs: ", err)
	}

	var devFound = false

	for _, dev := range ifs {
		if dev.Name == iface {
			devFound = true
		}
	}

	if !devFound && iface != "any" {
		log.Printf("error: did not find the interface %v\n", iface)
		return
	}

	// Get a BPF filter handle that we can set the filter on.
	handle, err := pcap.OpenLive(iface, snapLen, promisc, timeout)
	if err != nil {
		log.Printf("error: pcap.OpenLive failed: %v\n", err)
	}
	defer handle.Close()

	err = handle.SetBPFFilter(*filter)
	if err != nil {
		log.Printf("error: handle.SetBPFFilter failed: %v\n", err)
	}

	IPMap := map[string]map[string]data{}

	timeStart := time.Now()

	// gopacket.NetPacketSource will return a channel that we range over
	src := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range src.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		// If it is a real packet, check the content of the packet, and
		// update the map with the new values.
		if ipLayer != nil {
			createMapValue(ipLayer, packet, IPMap)
		}

		printMap(IPMap, timeStart)
	}
}
