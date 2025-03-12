package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	js, _ := jetstream.New(nc)

	ctx, cancel := context.WithTimeout(context.Background(), 999999*time.Second)
	defer cancel()

	kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: "keys",
	})

	name := <-js.StreamNames(ctx).Name()
	fmt.Printf("KV stream name: %s\n", name)

	go func() {
		for i := 0; i < 999999; i++ {
			time.Sleep(time.Second * 1)
			kv.Put(ctx, "node1.public", []byte(strconv.Itoa(i)))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
