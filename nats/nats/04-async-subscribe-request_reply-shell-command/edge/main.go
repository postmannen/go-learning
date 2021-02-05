package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

type Message struct {
	ID   int    // The Unique ID of the message
	Data string // The actual data in the message
}

func main() {
	// Create a connection to nats server, and publish a message.
	nc, err := nats.Connect("localhost", nil)
	if err != nil {
		log.Printf("error: nats.Connect failed: %v\n", err)
	}
	defer nc.Close()

	// Create a channel to put the data received in the subscriber callback
	// function
	reqMsgCh := make(chan Message)

	// Subscribe will start up a Go routine calling the callback function
	// specified when a new message is received.
	_, err = nc.Subscribe("subject1", func(req *nats.Msg) {
		message := Message{}

		// Create a buffer to decode the gob encoded binary data back
		// to it's original structure.
		buf := bytes.NewBuffer(req.Data)
		gobDec := gob.NewDecoder(buf)
		err := gobDec.Decode(&message)
		if err != nil {
			fmt.Printf("error: gob decoding failed: %v\n", err)
		}

		// Put the data recived on the channel for further processing
		reqMsgCh <- message

		// Send a confirmation message back to the publisher
		nc.Publish(req.Reply, []byte("confirmed: "+fmt.Sprint(message.ID)))
	})
	if err != nil {
		fmt.Printf("error: Subscribe failed: %v\n", err)
	}

	// Do some further processing of the actual data we received in the
	// subscriber callback function.
	for {
		fmt.Printf("subcriber: received data = %#v\n", <-reqMsgCh)
	}
}
