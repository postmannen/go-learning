/*
chan struct{} are used for signaling.
That means we don't necesarily put an empty struct{}
on that channel, but we can use it for controlling 
things in the program by closing the channel when we
are done.

In this example we use it to send a signal by closing
the channel after 3 seconds, and kind of making a timer.
The main program will block on <-c until channel is closed
or an struct{} is received, in our example here we close
the channel.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 3)
		close(c)
	}()

	<-c
	fmt.Println("Hello, playground")
}
