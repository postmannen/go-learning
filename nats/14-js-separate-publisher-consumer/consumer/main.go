package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	ctx := context.Background()

	nc, err := nats.Connect("localhost:4222", nats.Name("orders-consumer"))
	if err != nil {
		log.Fatalf("error: nats connect failed: %v\n", err)
	}
	defer nc.Close()

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf("error: jetstream new failed: %v\n", err)
	}

	stream, err := js.Stream(ctx, "orders")
	if err != nil {
		log.Fatalf("error: js.Stream failed: %v\n", err)
	}

	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:    "order_processor",
		Durable: "order_processor",
	})
	if err != nil {
		log.Fatalf("error: create or update consumer failed: %v\n", err)
	}

	cctx, err := consumer.Consume(func(msg jetstream.Msg) {
		log.Printf("Received message: %v\n", string(msg.Subject()))
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("error: create or update consumer failed: %v\n", err)
	}

	defer cctx.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
