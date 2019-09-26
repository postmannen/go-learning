package main

import (
	"fmt"
	"math/rand"
	"time"
)

// makeInt will return a random int
func makeInt() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(10) + 1
}

// makeIntWithTimer will return a random int with a little time delay
func makeIntWithTimer() int {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Second * 2)
	return rand.Intn(10) + 1
}

func printInts(a int, b int) {
	fmt.Println(a, b)
}

func main() {
	// Create two channels with a buffer of 1
	a := make(chan int, 1)
	b := make(chan int, 1)

	// Do the filling of the channels inside a goroutine,
	// so the code will continue with it's next step.
	go func() {
		a <- makeInt()
		b <- makeIntWithTimer()
	}()

	fmt.Println("a")

	startTime := time.Now()

	// The code will block here waiting for both a and b to be filled
	// before the printInts function is executed.
	printInts(<-a, <-b)

	endTime := time.Since(startTime)
	fmt.Println("The time it took = ", endTime)
}
