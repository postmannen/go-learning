package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Create a connection to nats server, and publish a message.
	nc, err := nats.Connect("localhost", nil)
	if err != nil {
		log.Printf("error: nats.Connect failed: %v\n", err)
	}
	defer nc.Close()

	// -----------Subscribe

	// Create a channel to put the data received in the subscriber callback
	// function
	subCh := make(chan []byte)

	_, err = nc.Subscribe("subject1", func(req *nats.Msg) {
		// Put the data recived on the channel for further processing
		subCh <- req.Data
		// Send a comfirmation message back to the publisher
		nc.Publish(req.Reply, []byte("confirmed"))
	})
	if err != nil {
		fmt.Printf("error: Subscribe failed: %v\n", err)
	}

	// Do some further processing of the actual data we received in the
	// subscriber callback function.
	for {
		fmt.Printf("subcriber: received data = %s\n", <-subCh)
	}
}
