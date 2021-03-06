package main

import (
	"fmt"
)

type holder struct {
	fu func(chan int, chan struct{})
}

func (h holder) Run(inCh chan int, done chan struct{}) {
	h.fu(inCh, done)
}

func main() {
	h := holder{
		fu: func(in chan int, done chan struct{}) {
			fmt.Println(<-in + 100)

			done <- struct{}{}
		},
	}

	valueCh := make(chan int)
	done := make(chan struct{})

	go h.Run(valueCh, done)

	valueCh <- 10
	<-done

}
