package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go out(ch1)

	for i := 1; i <= 5; i++ {
		ch1 <- i
	}

	time.Sleep(time.Second * 3)
	close(ch1)

}

func out(c1 chan int) {
	data, ok := <-c1
	if ok {
		fmt.Printf("Received on channel = %v, channel status = %v", data, ok)
	} else {
		fmt.Println("Channel closed")
	}
}
