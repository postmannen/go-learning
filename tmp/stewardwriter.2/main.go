package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	socketFullPath := flag.String("socketFullPath", "", "the full path to the steward socket file")
	messageFullPath := flag.String("messageFullPath", "", "the full path to the message")
	interval := flag.Int("interval", 10, "the interval in seconds between sending messages")
	flag.Parse()

	if *socketFullPath == "" {
		log.Printf("error: you need to specify the full path to the socket\n")
		return
	}
	if *messageFullPath == "" {
		log.Printf("error: you need to specify the full path to the message\n")
		return
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ticker := time.NewTicker(time.Second * time.Duration(*interval))

	for {
		select {
		case <-ticker.C:
			func() {
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

				log.Printf("info: succesfully wrote message to socket\n")
			}()
		case <-sigCh:
			log.Printf("info: received signal to quit..\n")
			return
		}
	}
}
