package main

import "fmt"
import bufchan "github.com/postmannen/go-learning/concurrency/17-slice-buffer-reading-from-channel"

func main() {
	//fake an input channel, and make it feed values to it.
	chIn := make(chan int)
	go func() {
		for i := 1; i <= 20; i++ {
			chIn <- i
		}
		close(chIn)
	}()

	b := bufchan.NewBuffer(5)
	b.Start(chIn)

	//loop and read a value from the out channel, and also show the content
	// of the buffer. The amout of values should be specified with the size
	// parameter to the buffer
	for v := range b.ChOut {
		fmt.Println("client : Read from chout : ", v)
		fmt.Println("client : Content of b.slice : ", b.Slice)
		b.ReadNext() //send msg to go routine to continue looping, and get another value from in channel.
	}
}
