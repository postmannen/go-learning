package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	socketFullPath := flag.String("socketFullPath", "", "the full path to the steward socket file")
	messageFullPath := flag.String("messageFullPath", "", "the full path to the message")
	flag.Parse()

	if *socketFullPath == "" {
		log.Printf("error: you need to specify the full path to the socket\n")
		return
	}
	if *messageFullPath == "" {
		log.Printf("error: you need to specify the full path to the message\n")
		return
	}

	socket, err := net.Dial("unix", *socketFullPath)
	if err != nil {
		log.Printf(" * failed: could not open socket file for writing: %v\n", err)
		return
	}
	defer socket.Close()

	fp, err := os.Open(*messageFullPath)
	if err != nil {
		log.Printf(" * failed: could not open message file for reading: %v\n", err)
		return
	}
	defer fp.Close()

	_, err = io.Copy(socket, fp)
	if err != nil {
		log.Printf("error: io.Copy failed: %v\n", err)
		return
	}
}
