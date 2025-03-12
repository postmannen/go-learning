package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	js, _ := jetstream.New(nc)

	cons, _ := js.CreateOrUpdateConsumer(ctx, "KV_keys", jetstream.ConsumerConfig{
		AckPolicy: jetstream.AckNonePolicy,
	})

	fmt.Printf("%+v\n", cons.CachedInfo())

	cctx, err := cons.Consume(func(msg jetstream.Msg) {
		log.Printf("\n message, subject: %v, data: %v\n", string(msg.Subject()), string(msg.Data()))
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
