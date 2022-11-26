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
				continue
			}

			counter++
			time.Sleep(time.Second * 1)
		}
	}()

	// -----------Subscribe

	go func() {
		_, err = nc.Subscribe("subject1", func(req *nats.Msg) {
			// Put the data recived on the channel for further processing
			fmt.Printf("subcriber1: received data = %s\n", req.Data)
		})
		if err != nil {
			fmt.Printf("error: Subscribe failed: %v\n", err)
		}
	}()

	go func() {
		_, err = nc.Subscribe("subject1", func(req *nats.Msg) {
			// Put the data recived on the channel for further processing
			fmt.Printf("subcriber2: received data = %s\n", req.Data)
		})
		if err != nil {
			fmt.Printf("error: Subscribe failed: %v\n", err)
		}
	}()

	select {}

}
