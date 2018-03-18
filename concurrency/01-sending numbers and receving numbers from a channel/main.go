package main

import (
	"fmt"
	"time"
)

func main() {
	//make a channel
	ch1 := make(chan int)
	defer close(ch1)

	go out(ch1)
	go in(ch1)

	time.Sleep(time.Second * 100)
}

func out(ch chan int) {
	for i := 1; i < 100; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}

}

func in(ch chan int) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}
