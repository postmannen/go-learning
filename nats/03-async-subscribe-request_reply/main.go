// Same as 02 example, but using PublishMsg method instead of Request method
// for publishing.

package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Create a connection to nats server, and publish a message.
	nc, err := nats.Connect("localhost", nil)
	if err != nil {
		log.Printf("error: nats.Connect failed: %v\n", err)
	}
	defer nc.Close()

	// ---------------Publish

	// The SubscribeSync used in the subscriber, will get messages that
	// are sent after it started subscribing, so we start a publisher
	// that sends out a message every second.
	go func() {
		counter := 0

		for {
			msg := &nats.Msg{
				Reply:   "subject2",
				Data:    []byte("just some data " + strconv.Itoa(counter)),
				Subject: "subject1",
			}

			err := nc.PublishMsg(msg)
			//msg, err := nc.Request("subject1", []byte("just some data "+strconv.Itoa(counter)), time.Second*2)
			if err != nil {
				log.Printf("error: publish failed: %v\n", err)
			}

			subRepl, err := nc.SubscribeSync(msg.Reply)
			if err != nil {
				log.Printf("error: nc.SubscribeSync failed: %v\n", err)
			}

			msgRepl, err := subRepl.NextMsg(time.Second * 10)
			if err != nil {
				log.Printf("error: subRepl.NextMsg failed: %v\n", err)
				// did not receive a repply, continuing from top again
			}
			fmt.Printf("publisher: received: %s\n", msgRepl.Data)

			counter++
			time.Sleep(time.Second * 1)
		}
	}()

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
