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

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	js, _ := jetstream.New(nc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: "profiles",
	})

	kv.Put(ctx, "sue.color", []byte("blue"))
	// entry, _ := kv.Get(ctx, "sue.color")
	// fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

	name := <-js.StreamNames(ctx).Name()
	fmt.Printf("KV stream name: %s\n", name)

	// -------------------------------

	go func() {
		nc, _ := nats.Connect(nats.DefaultURL)
		defer nc.Drain()

		js, _ := jetstream.New(nc)

		cons, _ := js.CreateOrUpdateConsumer(ctx, "KV_profiles", jetstream.ConsumerConfig{
			AckPolicy: jetstream.AckNonePolicy,
		})

		fmt.Printf("%+v\n", cons.CachedInfo())

		cctx, err := cons.Consume(func(msg jetstream.Msg) {
			log.Printf("\n *** Received message: %v, with data: %v\n", string(msg.Subject()), string(msg.Data()))
			msg.Ack()
		})
		if err != nil {
			log.Fatalf("error: create or update consumer failed: %v\n", err)
		}

		time.Sleep(time.Second * 5)

		defer cctx.Stop()
	}()

	time.Sleep(time.Second * 1)
	kv.Put(ctx, "sue.color", []byte("green"))
	time.Sleep(time.Second * 1)

}
