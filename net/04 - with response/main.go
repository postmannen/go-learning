package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//listen for inncomming connections
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error : Listener : ", err)
	}
	defer listener.Close()

	for {
		//Listen for new individual connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error : conn : ", err)
			os.Exit(1)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	//make a buffer to hold incomming data
	buf := make([]byte, 1024)

	//read the incomming connection into the buffer
	readLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error: Reading buffer : ", err, readLen)
	}

	//send response back to the network client
	conn.Write([]byte("Message received !"))

	conn.Close()
}
