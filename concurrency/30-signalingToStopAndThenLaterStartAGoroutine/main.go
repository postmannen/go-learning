package main

import (
	"fmt"
	"sync"
	"time"
)

func loopData(data []string, chStop chan struct{}, chStart chan struct{}, wg *sync.WaitGroup) {

	for _, word := range data {
		fmt.Println(word)
		time.Sleep(500 * time.Millisecond)

		// Check if we've got any signal for stopping the reading.
		select {
		// if a value received on chStop, it will block waiting
		// until a signal is received on the chStart channel.
		case <-chStop:
			<-chStart

		// if no value where received above, we do the default,
		// and continue with the next iteration of the for loop.
		default:
			continue
		}
	}
	wg.Done()

}

func main() {
	data := []string{"a", "b", "c", "d", "e", "f"}
	chStart := make(chan struct{})
	chStop := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)

	go loopData(data, chStop, chStart, &wg)

	time.Sleep(time.Second * 2)
	// signal that we want to halt the reading done in the go routine.
	chStop <- struct{}{}
	time.Sleep(time.Second * 2)
	// signal that we want to start the reading done in the go routine.
	chStart <- struct{}{}

	wg.Wait()
}
