package main

import (
	"fmt"
	"net"
)

func main() {
	//start a new network listener
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error: net.Listen: ", err)
	}

	defer listener.Close()

	for {
		//accept new connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: listener.Accept: ", err)
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Error: conn.Read: ", err)
			break
		}
		fmt.Println("bytes read = ", n)
		fmt.Println(string(data))

		//if the first 4 characters contain "exit", then break the for loop
		if string(data[:4]) == "exit" {
			fmt.Println("found exit")
			break
		}

	}
	fmt.Println("exiting handleFunc")
}
