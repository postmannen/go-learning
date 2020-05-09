package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {
	// Create a channel to keep track of the goroutines thats finnished.
	idCh := make(chan int)
	go func() {
		for {
			fmt.Printf("closed go routine %v\n", <-idCh)
		}
	}()

	// Start a network listener.
	l, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Printf("error: net.Listen: %v\n", err)
	}

	var connectionNR int

	// Wait for a new connection, and start doCommand for each of them.
	for {
		connectionNR++
		conn, err := l.Accept()
		if err != nil {
			log.Printf("error: l.Accept: %v\n", err)
		}

		go doCommand(conn, idCh, connectionNR)
	}

}

func doCommand(conn net.Conn, idCh chan int, id int) {
	for {
		// Create a scannner to read the input lines
		s := bufio.NewScanner(conn)

		// Prepare a bash command. The command to be
		// exexuted inside the bash will come from stdin.
		cmd := exec.Command("/bin/sh")
		// Since stdout wants an io.Writer we can pass conn directly to it.
		cmd.Stdout = conn
		cmd.Stderr = conn

		// Scan one line as string
		s.Scan()
		text := s.Text()
		if text == "exit" {
			conn.Close()
			idCh <- id
			return
		}

		// convert the stdin command to a reader, and pass to stdin.
		cmd.Stdin = strings.NewReader(text)
		err := cmd.Run()
		if err != nil {
			log.Printf("error: cmd.Run: %+v\n", err)
		}
	}

}
