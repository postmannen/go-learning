package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

func main() {
	getPortList()
	fmt.Println("--------------------------------------")
	getPortListDetails()
}

func getPortListDetails() {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
		return
	}
	for _, port := range ports {
		fmt.Printf("Found port: %s\n", port.Name)
		if port.IsUSB {
			fmt.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
			fmt.Printf("   USB serial %s\n", port.SerialNumber)
		}
	}
}

func getPortList() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}
