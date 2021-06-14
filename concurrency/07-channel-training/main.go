package main

import (
	"fmt"
)

// Set how many times something should be put on the channel.
// Put on the channel, and close channel when last value is put.
// Read all values on the channel, and exit when last is read.
//
func main() {
	numberOfTimes := 10
	ch := make(chan int, 1)

	go putOnChannelNumberOfTimes(numberOfTimes, ch)

	for v := range ch {
		fmt.Println("number = ", v)
	}

}

func putOnChannelNumberOfTimes(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	fmt.Println("closing channel")
	close(c)

}
