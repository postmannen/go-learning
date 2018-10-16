/*
The idea here is to simulate processing data and logging
of eacg process of data.
The processing of data and logging is split into two workers,
where each have it's own queue.
Each time new data is placed on the pipeline, the same data
will be put on both the data channel and the log channel. But
if the log channel is full, a message will be printed in the
console.
On purpose the routine that reads logs is given a delay, so it
will not work as fast as the one who handles data, to simulate
hickup's in the logging system.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type testData struct {
	data    string
	log     string
	counter int
}

var mu sync.Mutex

func main() {
	chData := make(chan testData, 10)
	chLog := make(chan testData, 10)
	counter := 0

	//Create some data, and some logs
	//Will create 'i' nr of Go routines to push data
	for i := 0; i < 10; i++ {
		go createData(i, &counter, chData, chLog)
	}

	//Process data
	go func() {
		for {
			d := <-chData
			fmt.Println("Processing data : ", d)
		}
	}()

	//Process logs

	for ii := 0; ii < 1000; ii++ {
		//Putting in a time.Sleep here to simulate a slow worker
		time.Sleep(time.Millisecond * 550)
		l := <-chLog
		fmt.Println("###Processing log : ", l)
	}

}

//createData goNR to keep track of Go routine,
//counter the batch of data,
//chData the data we prioritize,
//chLog the log data that can be discarded if slow system,
func createData(goNR int, counter *int, chData chan testData, chLog chan testData) {
	for {
		d := testData{
			data: fmt.Sprint("some data", "GR=", goNR, ","),
			log:  fmt.Sprint("some log", "GR=", goNR, ","),
		}

		time.Sleep(time.Millisecond * 100)

		//Counter is updated from multiple Go routines, so we need a mutex
		mu.Lock()
		*counter++
		mu.Unlock()
		d.counter = *counter

		//Always let the real data put data on the channel
		chData <- d

		//If the log channel is full, don't put it on the channel
		select {
		case chLog <- d: //Put on channel if room
		default: //If not room
			fmt.Printf("***Log channel full for counter %v ***\n", *counter)
		}

	}
}
