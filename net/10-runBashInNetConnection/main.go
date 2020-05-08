package main

import (
	"bufio"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {
	// Start a network listener.
	l, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Printf("error: net.Listen: %v\n", err)
	}

	//stopCh := make(chan chan struct{}, 1)

	// Wait for a new connection, and start doCommand for each of them.
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("error: l.Accept: %v\n", err)
		}

		go doCommand(conn)
	}

}

func doCommand(conn net.Conn) {
	for {
		// Create a scannner to read the input lines
		s := bufio.NewScanner(conn)

		// Prepare a bash command. The command to be
		// exexuted inside the bash will come from stdin.
		cmd := exec.Command("/bin/bash")
		// Since stdout wants an io.Writer we can pass conn directly to it.
		cmd.Stdout = conn

		// Scan one line as string
		s.Scan()
		text := s.Text()
		if text == "exit" {
			return
		}

		// convert the stdin command to a reader, and pass to stdin.
		cmd.Stdin = strings.NewReader(text)
		err := cmd.Run()
		if err != nil {
			log.Printf("error: cmd.Run: %v\n", err)
			return
		}
	}

}
