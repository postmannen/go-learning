package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	ctx := context.Background()

	nc, err := nats.Connect("localhost:4222", nats.Name("order-publisher"))
	if err != nil {
		log.Fatalf("error: nats connect failed: %v\n", err)
	}
	defer nc.Close()

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf("error: jetstream new failed: %v\n", err)
	}

	_, err = js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:        "orders",
		Description: "orders stream",
		Subjects:    []string{"orders.>"},
	})

	if err != nil {
		log.Fatalf("error: jetstream create or update failed: %v\n", err)
	}

	i := 0
	for {
		_, err := js.Publish(ctx, fmt.Sprintf("orders.%v", i), []byte("hello world"))
		if err != nil {
			log.Fatalf("error: js failed to publish message: %v\n", err)
		}

		log.Printf("published message: %v\n", i)
		time.Sleep(time.Second * 1)
		i++
	}
}
