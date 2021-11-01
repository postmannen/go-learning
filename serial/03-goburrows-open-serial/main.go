package main

import (
	"fmt"
	"log"

	"github.com/goburrow/serial"
	"go.bug.st/serial/enumerator"
)

func main() {
	getPortListDetails()
	fmt.Println("-------------------------------------------")

	conf := &serial.Config{Address: "/dev/ttyUSB0",
		BaudRate: 9600,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
	}

	port, err := serial.Open(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	fmt.Println(" * DEBUG * Before write")

	_, err = port.Write([]byte("serial\n\r"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(" * DEBUG * After write")

	for {
		b := make([]byte, 64)
		n, err := port.Read(b)
		if err != nil {
			log.Printf("error: port.Read: %v\n", err)
		}
		fmt.Printf("read %v number of characters\n", n)
	}

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
