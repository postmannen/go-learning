package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	iface   = "any"
	snapLen = int32(1600)
	promisc = false
	timeout = pcap.BlockForever
)

func main() {
	port := flag.String("port", "", "the port number to listen to, eg. -port 502")
	flag.Parse()

	if *port == "" {
		log.Println("error: you need to provide a port with the -port flag")
	}

	filter := "port " + *port

	if os.Getgid() != 0 {
		log.Println("error: Need root privileges")
		return
	}

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

	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Printf("error: handle.SetBPFFilter failed: %v\n", err)
	}

	// gopacket.NetPacketSource will return a channel that we range over
	src := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range src.Packets() {
		// fmt.Printf("%+v\n", packet)

		appLayer := packet.ApplicationLayer()
		if appLayer != nil {
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer != nil {
				ip, _ := ipLayer.(*layers.IPv4)
				fmt.Printf("src:%v -> dts:%v\n ", ip.SrcIP, ip.DstIP)
			}

			fmt.Printf("content: %v\n", string(appLayer.Payload()))
		}
	}
}
