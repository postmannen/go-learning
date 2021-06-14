package main

import (
	"fmt"
	"net"
	"time"
)

func netListen() {
	server, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Error: net.Listen: ", err)
	}
	time.Sleep(100000 * time.Millisecond)
	defer server.Close()
}

func main() {
	server, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Error: net.Listen: ", err)
	}
	defer server.Close()

	time.Sleep(9000000 * time.Millisecond)

}
