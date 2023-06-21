package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatalf("error: nats connection failed: %v\n", err)
	}

	defer nc.Drain()

	const streamName = "EVENTS"

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("error: jetstream failed: %v\n", err)
	}

	js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{"events.>"},
	})

	consumerName := "consumer1"
	ackWait := 10 * time.Second
	ackPolicy := nats.AckExplicitPolicy
	maxWaiting := 1

	_, err = js.AddConsumer(streamName, &nats.ConsumerConfig{
		Durable:       consumerName,
		AckPolicy:     ackPolicy,
		AckWait:       ackWait,
		MaxAckPending: maxWaiting,
	})

	if err != nil {
		log.Fatalf("error: add consumer failed: %v\n", err)
	}

	sub, _ := js.PullSubscribe("", consumerName, nats.Bind(streamName, consumerName))

	js.Publish("events.1", nil)
	js.Publish("events.2", nil)

	msgs, err := sub.Fetch(3)
	if err != nil {
		log.Fatalf("error: sub.Fetch 1 failed: %v\n", err)
	}
	fmt.Printf("requested 3, got %d\n", len(msgs))

	msgs[0].Ack()

	msgs, _ = sub.Fetch(1)
	fmt.Printf("requested 1, got %d\n", len(msgs))
	msgs[0].Ack()

}
