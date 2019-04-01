package main

import (
	"context"
	"fmt"
	"time"
)

//count takes a contect, and a channel to put numbers on as input.
func count(ctx context.Context, dst chan<- int) {
	n := 1
	for {
		select {
		case dst <- n:
			n++
		case <-ctx.Done():
			fmt.Println("not leaking")
			return
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ints := make(chan int)
	go count(ctx, ints)
	for n := range ints {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}

	time.Sleep(time.Second * 2)
}
