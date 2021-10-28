package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func main() {
	getPortList()
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
