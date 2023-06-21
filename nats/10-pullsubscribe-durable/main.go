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

	streamName := "EVENTS2"

	si, err := js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{"*.events.>"},
	})
	if err != nil {
		fmt.Printf("add stream failed: %v\n", err)
	}

	fmt.Printf("stream info: %v\n", si)

	_, err = js.Publish("no.events.node1", []byte("event1"))
	if err != nil {
		fmt.Printf("publish failed: %v\n", err)
	}
	_, err = js.Publish("se.events.node2", []byte("event2"))
	if err != nil {
		fmt.Printf("publish failed: %v\n", err)
	}
	_, err = js.Publish("gb.events.node3", []byte("event3"))
	if err != nil {
		fmt.Printf("publish failed: %v\n", err)
	}

	// Create a subscription.
	sub, _ := js.PullSubscribe("", "node", nats.BindStream(streamName))

	for i := 0; i < 3; i++ {
		msgs, err := sub.Fetch(1)
		if err != nil {
			fmt.Printf("fetch failed: %v\n", err)
		}

		for i, v := range msgs {
			fmt.Printf("i=%v, data: %s,subject: %s\n", i, v.Data, v.Subject)
			//fmt.Printf("%#v\n", v)
			err := v.Ack()
			if err != nil {
				fmt.Printf("fetch failed: %v\n", err)
			}
		}
	}

	for i := 0; i < 3; i++ {
		msgs, err := sub.Fetch(1)
		if err != nil {
			fmt.Printf("fetch failed: %v\n", err)
		}

		for i, v := range msgs {
			fmt.Printf("i=%v, data: %s,subject: %s\n", i, v.Data, v.Subject)
			//fmt.Printf("%#v\n", v)
			err := v.Ack()
			if err != nil {
				fmt.Printf("fetch failed: %v\n", err)
			}
		}
	}

	js.PurgeStream(streamName)
	sub.Drain()

}
