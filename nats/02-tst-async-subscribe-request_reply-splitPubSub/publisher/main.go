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
	nc, err := nats.Connect("localhost", nats.Timeout(time.Second*5))
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
			b := []byte("just some data " + strconv.Itoa(counter))

			msg, err := nc.Request("subject1", b, time.Second*20)
			if err != nil {
				log.Printf("error: publish failed: %v\n", err)
				continue
			}

			fmt.Printf("publisher: received: %s\n", msg.Data)

			counter++
			time.Sleep(time.Second * 1)
		}
	}()

	select {}

}
