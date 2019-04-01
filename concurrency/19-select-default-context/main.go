package main

import (
	"context"
	"fmt"
	"time"
)

func a(ctx context.Context) {
	go b(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("received ctx.Done in A, breaking out")
			break
		default:

		}

		fmt.Println("printing a ........")
		time.Sleep(time.Millisecond * 200)
	}
}

func b(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("received ctx.Done in B, breaking out")
			break
		default:

		}

		fmt.Println("printing b ........")
		time.Sleep(time.Millisecond * 200)
	}
}
func main() {
	ctx := context.Background()

	go a(ctx)
	time.Sleep(time.Second * 1)
	ctx.Done()
	fmt.Println("done !")
}
