package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to the NATS server.
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create JetStream context.
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	// Create a stream.
	_, err = js.AddStream(&nats.StreamConfig{Name: "TEST_STREAM", Subjects: []string{"orders.*"}})
	if err != nil {
		log.Fatal(err)
	}

	// Publish a message to the stream.
	for i := 0; i < 20; i++ {
		_, err = js.Publish(fmt.Sprintf("orders.%v", i), []byte(fmt.Sprintf("Order %v", i)))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create a consumer.
	_, err = js.AddConsumer("TEST_STREAM", &nats.ConsumerConfig{Durable: "orders-consumer", AckPolicy: nats.AckExplicitPolicy})
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to the consumer.
	sub, _ := js.PullSubscribe("orders.*", "orders-consumer")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ----------------------------

	streamInfo, err := js.StreamInfo("TEST_STREAM", nats.Context(ctx))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of messages in the stream: %d\n", streamInfo.State.Msgs)

	// -----------------------------

	// Fetch and process messages.
	for {
		msgs, err := sub.Fetch(10, nats.Context(ctx))
		if err != nil {
			log.Fatal(err)
		}

		for _, msg := range msgs {
			fmt.Printf("Received message: %s\n", string(msg.Data))
			// Acknowledge the message.
			msg.Ack()
		}
	}
}
