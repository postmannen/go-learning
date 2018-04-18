package main

import (
	"fmt"
)

func main() {
	numberOfTimes := 10
	ch := make(chan int, 1)

	go printNumberOfTimes(numberOfTimes, ch)

	for v := range ch {
		fmt.Println("number = ", v)
	}

}

func printNumberOfTimes(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	fmt.Println("closing channel")
	close(c)

}
