package main

import "fmt"

// ringBuffer holds the data of the buffer,
type ringBuffer struct {
	data chan string
}

// newringBuffer is a push/pop storage for values.
func newringBuffer() *ringBuffer {
	return &ringBuffer{
		data: make(chan string, 10),
	}
}

// startWithChannels
func (s *ringBuffer) startWithChannels(inCh chan string, outCh chan string) {
	// Fill the buffer when new data arrives
	go func() {
		for v := range inCh {
			s.data <- v
			fmt.Printf("**BUFFER** DEBUG PUSHED ON BUFFER: value = %v\n\n", v)
		}
		close(s.data)
	}()

	go func() {
		for v := range s.data {
			outCh <- v
		}

		close(outCh)

	}()
}

func main() {
	rb := newringBuffer()
	inCh := make(chan string)
	outCh := make(chan string)

	rb.startWithChannels(inCh, outCh)

	inCh <- "apekatt"
	inCh <- "hest"
	inCh <- "ku"
	close(inCh)

	for v := range outCh {
		fmt.Printf("got: %v\n", v)
	}

}

