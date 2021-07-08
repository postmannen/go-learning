// Signaling with CallBack.
// Embedding the signal send to channel inside a callback function.
// The signal will be received in main when the done func is called,
// and we can continue the work.

package main

import "fmt"

func doSome(done func()) {
	fmt.Printf("Doing some work\n")
	done()
}

func main() {
	doneCh := make(chan struct{})

	done := func() {
		doneCh <- struct{}{}
	}

	for i := 0; i < 1000; i++ {
		go doSome(done)
		<-doneCh
	}

}
