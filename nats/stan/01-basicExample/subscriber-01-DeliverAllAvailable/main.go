package main

import (
	"fmt"
	"log"

	stan "github.com/nats-io/stan.go"
)

func main() {

	sc, err := stan.Connect("test-cluster", "clientID2")
	if err != nil {
		log.Printf("error: stan.Connect failed: %v\n", err)
	}

	// Create a function that will be the actual callback function
	// that will be called each time a new message is received.
	// By default Stan ack's messages, but if manual ack mode was
	// set we would add it here with msg.Ack().
	subscribeFunc := func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}

	// Subscribe is aSync or callback.
	// When we call the sc.Subscribe method we actually start a Go routine
	// behind the scenes that will listen for received messages, and call
	// the specified subscribeFunc for each message received.
	sub, err := sc.Subscribe("foo", subscribeFunc, stan.DeliverAllAvailable())
	if err != nil {
		log.Printf("error: sc.Subscribe failed: %v\n", err)
	}

	// Unsubscribe
	defer sub.Unsubscribe()

	// Close connection
	defer sc.Close()

	select {}
}
