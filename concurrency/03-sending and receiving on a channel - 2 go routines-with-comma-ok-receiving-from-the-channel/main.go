package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// Creating a channel without a buffer.
	// This will make the calling go routine to wait until a recevier is up and processing the data.
	// If we wanted a buffer we could specify it this way 'ch1 := make(chan int,5)' which will make a channel with a buffer of 5.
	// With a buffer of 5 the calling go routine will be able to put 5 int's in the channel,
	// before it have to wait for the receiver to process the data on the channel
	ch1 := make(chan int)

	// no wait group adding here, since the 'get' function will not be finnished before 'put' is done, and the channel is closed either way.
	go put(ch1)

	// Adding 1 to the waitgroup since there is only 1 go routine we will wait on
	wg.Add(1)
	go get(ch1)

	// Will wait here until the go routine are finnished and the waitgroup is zero before it continues
	wg.Wait()
}

func put(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
	}
	// Closing the channel on the sender when the job is done.
	// NB : The sender shall allways close the channel, never the receiver.
	close(c)
}

func get(c chan int) {
	for {
		// Using two variables with channels, the 2nd variable will have the state of the channel as type bool
		myVar, ok := <-c
		fmt.Printf("The value of 'ok' = %v, and the type is %T\n", ok, ok)
		if ok {
			fmt.Println("content of myVar = ", myVar, "and the content of c = ", c)
			time.Sleep(time.Millisecond * 500)
		} else {
			// The channel is closed since 'ok=false', and we do the logic to decrement the waitgroup, and leave the function
			fmt.Println("Channel was closed, returning for loop")
			// Decrement the waitgroup value by 1 by calling wg.Done()
			wg.Done()
			return
		}
	}
}
