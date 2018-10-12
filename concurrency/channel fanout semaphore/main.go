/*
Channel Fanout Semaphore.
The idea is here is to do a fanout which spawns many Go routines
writing they're result to 1 channel,and have 1 worker that
consume the data they produce.
To limit the number of Go routines posting to the channel
at the same time we can create a semaphore (semaphoreNR), which is
another channel with the len set to the number of Go
routines we will allow to do work at the same time. When
each individual Go routine is done the value put on the
channel to block it will be released, and a new Go routine
can take it's place and start working.
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

type values struct {
	text        string
	semaphoreNR int
}

const goRoutines = 100
const semaphore = 10

func main() {
	runtime.GOMAXPROCS(1)

	ch := make(chan values, goRoutines)
	sem := make(chan bool, semaphore)

	for i := 0; i < goRoutines; i++ {
		//create a literal (unnamed) function to spawn a Go routine
		go func() {
			//Put a value on the limited buffer sem channel.
			//If the channel buffer is full, the code will block here
			//until the a spot in the channel buffer is made free, and
			//we can take it's place by putting a value on the channel,
			//and continue with the rest of the code.
			//
			sem <- true

			//putting in a little delay so it will spend some time
			//with each spawned Go routine, or else it will be to
			//quick and the CPU can handle it without the need of
			//concurrency.
			//
			time.Sleep(time.Nanosecond * 1000)
			ch <- values{
				text:        "go routines running",
				semaphoreNR: len(sem), //Number of Go routines running now
			}

			//Free up a spot on the buffered sem channel for a new
			//Go routine to take it's place.
			//
			<-sem
		}()
	}

	//Loop through nr. of goroutines, and extract the data from the main channel.
	for i := goRoutines; i > 0; i-- {
		c := <-ch
		fmt.Println(c)
	}

}
