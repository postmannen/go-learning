package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	//start a litstening tcp server
	server, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Error: net.Listen: ", err)
	}
	defer server.Close()

	//create an endless loop
	for {
		//create a new connection/session if one is connected
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error: creating conn :")
		}

		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)

		//time.Sleep(9000000 * time.Millisecond)
	}
}
