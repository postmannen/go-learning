/*
Starts a tcp listener on localhost port 9000,
for testing net.Conn, and do something based on the input
*/
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

	for {
		myC := string("Enter command : ")
		myCommand := []byte(myC)
		_, err := conn.Write(myCommand)
		if err != nil {
			fmt.Println("error: failed to print command to session")
		}

		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error: reading Conn to buf: ", err)
			break
		}

		//convert buffer to string
		s := string(buf)

		if s[:3] == "ape" {
			fmt.Println("du skrev ape du !")
		}

		fmt.Println("Reading ", n, "bytes, which contains the string : ", s)

	}

}
