/*
 Spawn up a bunch of goroutines who want to put a single nr. on a channel.
 Read that channel from a single go routine.
 The time.Sleep is for seing in the console that all go routines have been
 created and is present before we start reading.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

const nrGoRoutines int = 1000

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 0; i < nrGoRoutines; i++ {
			//Start one go routine for each value, so there will be as many go routines
			// waiting to put someting on the channel as there are values.
			wg.Add(1)
			go func(ii int) {
				fmt.Printf("Waiting to put  %v on channel\n", ii)
				c <- ii
				fmt.Printf("put %v int on channel\n", ii)
				wg.Done()
			}(i)
		}
		//close(c) //Putting the close here will panic with "trying to send on a close channel".
		wg.Done()
	}()

	//The time.Sleep is for seing in the console that all go routines have been
	//created and is present before we start reading.
	time.Sleep(time.Second * 1)

	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	wg.Wait()
	close(c)
}
