//Example for checking if a channel is closed.
//
//A channel actually carries two values. One for the value itself, and what if the channel is closed
// or not. When a channel is close it will return 0 for the value received, and false on OK, which
// is the second value to be received from a channel.
package main

import (
	"fmt"
)

//aFunc takes a channel as input, and will put 9 values on the channel before closing it.
func aFunc(c chan int) {
	for i := 1; i < 10; i++ {
		c <- i
	}
	close(c)

	//Uncommenting the line below will show error sending to closed channel,
	//even though the channel got a buffer of 100
	//c <- 1
}

func main() {
	//Create a channel that got a buffer way higher than we need, just for showing.
	aCh := make(chan int, 100)
	aFunc(aCh)

	//Read the channel, and check if the channel is closed, exit if closed.
	for {
		v, ok := <-aCh
		fmt.Println(v)

		// if the channel is closed
		if v == 0 && !ok {
			fmt.Printf("The channel is closed, v = %v, and ok = %v\n", v, ok)
			fmt.Println("Exiting for loop")
			break
		}
	}

}
