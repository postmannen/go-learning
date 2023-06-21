package main

import (
	"fmt"
	"net"
	"os"

	"log"
)

const socketFile = "./a.sock"

func main() {
	// Delete any existing sockets
	err := os.Remove(socketFile)
	if err != nil {
		log.Printf("error: removing socket file: %v\n", err)
	}

	// The listener will will also create the sock file.
	l, err := net.Listen("unix", socketFile)
	if err != nil {
		log.Printf("error: net.Dial: %v\n", err)
		return
	}
	defer l.Close()

	// Accept connections, and read the data comming in.
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Printf("error: l.Accept: %v\n", err)
				return
			}

			go func(conn net.Conn) {
				b := make([]byte, 1024)
				_, err = conn.Read(b)
				if err != nil {
					log.Printf("error: failed to read bytes: %v\n", err)
				}

				fmt.Printf("read: %s\n", b)

				conn.Close()
			}(conn)
		}
	}()

	// Dial the socket, and write some data to it.
	conn, err := net.Dial("unix", socketFile)
	if err != nil {
		log.Printf("error: net.Dial failed: %v\n", err)
	}
	defer conn.Close()
	conn.Write([]byte("a horse jumped over the hill"))

	{
		conn, err := net.Dial("unix", socketFile)
		if err != nil {
			log.Printf("error: net.Dial failed: %v\n", err)
		}
		defer conn.Close()
		conn.Write([]byte("a horse jumped over the hill"))
	}

	{
		conn, err := net.Dial("unix", socketFile)
		if err != nil {
			log.Printf("error: net.Dial failed: %v\n", err)
		}
		defer conn.Close()
		conn.Write([]byte("a horse jumped over the hill"))
	}

	select {}

}
