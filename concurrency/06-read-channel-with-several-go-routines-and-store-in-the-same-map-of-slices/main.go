package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	// map to hold the numbers processed by each go routine
	mapOfSlices := make(map[int][]int)

	// Make a channel with a buffer of 5.
	ch1 := make(chan int, 5)

	// Tell the program that we will start 5 Go Routines,
	// so we know how many we shall wait for to finish.
	goRoutines := 5
	wg.Add(goRoutines)

	go createNumber(ch1)

	// Spawn 5 Go Routines who will try to read from the channel.
	for i := 1; i < goRoutines; i++ {
		go readNumber(ch1, i, mapOfSlices)
	}

	// Wait here until all Go Routines are done, when all done, continue.
	// Continue will happen when wg == 0, and all wg.Done() commands have
	// subtracted one value from the wg who was initially set to 5.
	wg.Wait()

	//Iterate the map of slices, and print out info about the numbers collected
	for k1, v1 := range mapOfSlices {
		fmt.Printf("--------Numbers collected for index %v = %v -----------\n", k1, len(v1))
		for _, v2 := range v1 {
			fmt.Printf("Index=%v, number=%v \n", k1, v2)
		}
	}

	fmt.Println(mapOfSlices)
}

//create sequential numbers, and put them on the channel
func readNumber(c1 chan int, funcIndex int, m map[int][]int) {
	for {
		v, ok := <-c1
		if ok {
			//need mutex here, since several go routines try to write to the map at the same time
			mu.Lock()
			m[funcIndex] = append(m[funcIndex], v)
			mu.Unlock()
		} else {
			break
		}
	}
	wg.Done()
}

func createNumber(c1 chan int) {
	i := 1
	for {
		c1 <- i
		i++
		if i > 50 {
			// Close the channel when we've reached the disered number.
			close(c1)
			// Break out of the for loop.
			break
		}
	}

	// Tell the workgroup that we are done with 1 Goroutine.
	// wg.Done() will subtract 1 number from the wg, and then
	// the wg.Wait command will know that one is finnished.
	wg.Done()
}
