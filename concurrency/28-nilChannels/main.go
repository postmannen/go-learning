// Example is for testing out setting channels to nil.
// Starting up 2 go routines who will produce a set of
// numbers, and put them on a channel for each go routine.
// When to go routine is done we close the channel.
//
// We then make a loop with a select, and read all the values
// from the 2 channels. With OK we check if the channel is
// closed, and if it is closed we set the channel to nil.
// The interesting thing bout settign a channel to nil is that
// it will not be checked anymore in the select, and the case
// is ignored. Something that can be seen in this example that
// the print statement inside the case about setting the channel
// to nil is only printed once.
//
// We then check that both channels are set to nil, and break
// out of the loop.
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i <= 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				fmt.Printf("setting ch1 to nil\n")
				ch1 = nil
			} else {
				fmt.Printf("* ch1: %v\n", v)
			}

		case v, ok := <-ch2:
			if !ok {
				fmt.Printf("setting ch2 to nil\n")
				ch2 = nil
			} else {
				fmt.Printf("** ch2: %v\n", v)
			}

		}

		if ch1 == nil && ch2 == nil {
			fmt.Printf("info: both channels closed, returning...\n")
			break
		}
	}

}
