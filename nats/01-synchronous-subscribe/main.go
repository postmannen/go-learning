package main

import (
	"fmt"
	"log"
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

	// Publish something
	// The SubscribeSync used in the subscriber, will get messages that
	// are sent after it started subscribing, so we start a publisher
	// that sends out a message every second.
	go func() {
		for {
			err := nc.Publish("subject1", []byte("A little test"))
			if err != nil {
				log.Printf("error: publish failed: %v\n", err)
			}

			time.Sleep(time.Second * 1)
		}
	}()

	// Subscribe
	sub, err := nc.SubscribeSync("subject1")
	if err != nil {
		fmt.Printf("error: SubscribeSync failed: %v\n", err)
	}

	for {
		msg, err := sub.NextMsg(time.Second * 5)
		if err != nil {
			fmt.Printf("error: sub.NextMsg failed: %v\n", err)
		}

		fmt.Printf("subcribed data = %s\n", msg.Data)
	}

}
