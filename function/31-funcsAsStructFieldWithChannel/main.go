package main

import (
	"fmt"
)

type holder struct {
	fu func(chan int, chan struct{})
}

func newHolder(f func(chan int, chan struct{})) holder {
	return holder{
		fu: f,
	}
}

func (h holder) Run(inCh chan int, done chan struct{}) {
	h.fu(inCh, done)
}

func main() {
	f := func(in chan int, done chan struct{}) {
		fmt.Println(<-in + 100)

		done <- struct{}{}
	}

	h := newHolder(f)

	valueCh := make(chan int)
	done := make(chan struct{})

	go h.Run(valueCh, done)

	valueCh <- 10
	<-done

}
