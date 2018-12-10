/*
This example is the same as 04, just that the checking of data is done with a range loop instead of checking if the loop is closed or not.
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//create a channel without a buffer
	ch1 := make(chan int)

	wg.Add(1)
	go get(ch1)

	go put(ch1)

	wg.Wait()
}

func put(c1 chan<- int) {
	for i := 1; i <= 5; i++ {
		c1 <- i
	}
	//closing the channel when done
	close(c1)
}

func get(c1 <-chan int) {
	for v := range c1 {
		fmt.Println("func 'get' received = ", v)
	}
	wg.Done()
}
