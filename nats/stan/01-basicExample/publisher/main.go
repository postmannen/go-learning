package main

import (
	"log"

	stan "github.com/nats-io/stan.go"
)

func main() {

	sc, err := stan.Connect("test-cluster", "clientID1")
	if err != nil {
		log.Printf("error: stan.Connect failed: %v\n", err)
	}

	// Simple Synchronous Publisher
	// does not return until an ack has been received from NATS Streaming
	err = sc.Publish("foo", []byte("Hello World2"))
	if err != nil {
		log.Printf("error: sc.Publish failed: %v\n", err)
	}

	err = sc.Publish("foo", []byte("Hello "))
	if err != nil {
		log.Printf("error: sc.Publish failed: %v\n", err)
	}

	// Close connection
	sc.Close()
}
