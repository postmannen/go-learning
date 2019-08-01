package main

import (
	"fmt"
	"time"
)

func main() {
	// Make an unbuffered channel. This means we create a channel that
	// can hold only one value at a time, so it have to be read from
	// before we can put another one in.
	// The channel will block a put call if it is full, and it will
	// also block a get call and wait if there is not anything yet
	// to receive on the channel.
	ch1 := make(chan int)
	defer close(ch1)

	// Start the functions as Go Routines.
	go out(ch1)
	go in(ch1)

	// A small sleep here to let the Go Routines finnish before we exit main().
	time.Sleep(time.Second * 100)
}

// out will have a for loop that creates values from 0-100, and puts those
// values on a channel one by one.
func out(ch chan int) {
	for i := 1; i < 100; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}

}

// in will try do read all the values comming in on a channel, and put them
// one by one in msg.
func in(ch chan int) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}
