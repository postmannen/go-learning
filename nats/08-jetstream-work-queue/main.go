package main

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	defer nc.Drain()

	js, _ := nc.JetStream()

	cfg := &nats.StreamConfig{
		Name:      "EVENTS",
		Retention: nats.WorkQueuePolicy,
		Subjects:  []string{"events.>"},
	}

	js.AddStream(cfg)
	fmt.Println("created the stream")

	fmt.Println("\n# Create non-overlapping consumers")
	sub1, _ := js.PullSubscribe("events.us.>", "processor-us", nats.BindStream(cfg.Name))
	sub2, _ := js.PullSubscribe("events.eu.>", "processor-eu", nats.BindStream(cfg.Name))

	js.Publish("events.eu.mouse_clicked", nil)
	js.Publish("events.us.page_loaded", nil)
	js.Publish("events.us.input_focused", nil)
	js.Publish("events.eu.page_loaded", nil)
	fmt.Println("published 4 messages")

	msgs, _ := sub1.Fetch(2)
	for _, msg := range msgs {
		fmt.Printf("us sub got: %s\n", msg.Subject)
		msg.Ack()
	}

	msgs, _ = sub2.Fetch(2)
	for _, msg := range msgs {
		fmt.Printf("eu sub got: %s\n", msg.Subject)
		msg.Ack()
	}
}
