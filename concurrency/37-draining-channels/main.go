package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a buffered channel and fill it with 10 values.
	ch := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)

	// A single go routine who read just on value from the channel.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Printf("%v\n", <-ch)
		wg.Done()
	}()
	wg.Wait()

	// For some odd reason we exit, and we want to drain the remainding
	// elemenets in the channel.
	// if some error..then..
	for range ch {
	}

	fmt.Printf("Elements in channel: %v\n", len(ch))
}
