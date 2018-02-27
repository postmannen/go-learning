package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	fmt.Println("Starting tcp listener at port 9000")
	if err != nil {
		fmt.Println("error: listener:", err)
		listener.Close()
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error: conn Accept:", err)
			conn.Close()
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error: reading Conn to buf: ", err)
		}

		fmt.Println("Reading ", n, "bytes")
	}

}
