package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ifs, err := net.Interfaces()
	if err != nil {
		log.Printf("error: Intefaces: %v\n", err)
		return
	}

	for _, i := range ifs {
		addr, _ := i.Addrs()
		for _, a := range addr {
			fmt.Printf("ifs : %v, addr: %v \n", i, a.String())
		}
	}

}
