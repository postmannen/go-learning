package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
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
		Bucket: "keys",
	})

	kv.Put(ctx, "node1.public", []byte("blue"))
	// entry, _ := kv.Get(ctx, "sue.color")
	// fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

	name := <-js.StreamNames(ctx).Name()
	fmt.Printf("KV stream name: %s\n", name)

	time.Sleep(time.Second * 1)
	kv.Put(ctx, "node1.public", []byte("green"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
