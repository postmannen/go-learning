package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	ifs, err := pcap.FindAllDevs()
	if err != nil {
		log.Println("error: pcap.FindAllDevs: ", err)
	}

	for _, dev := range ifs {
		fmt.Println("-----------------------")
		fmt.Printf("if: %v\n", dev.Name)
		for _, ip := range dev.Addresses {
			fmt.Printf("    ip: %v, mask: %v\n", ip.IP, ip.Netmask)
		}
	}
}
