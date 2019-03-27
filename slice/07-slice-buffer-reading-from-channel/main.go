/*
 Make a buffered reader of channel.
 Will keep the next input values read in a buffer where size if defined by b.size.
 Will release a new value with the readNext method.
*/
package main

import (
	"fmt"
)

func main() {
	//fake an input channel, and make it feed values to it.
	chIn := make(chan int)
	go func() {
		for i := 1; i <= 20; i++ {
			chIn <- i
		}
		close(chIn)
	}()

	b := NewBuffer(5)
	b.start(chIn)

	//loop and read a value from the out channel, and also show the content
	// of the buffer. The amout of values should be specified with the size
	// parameter to the buffer
	for v := range b.chOut {
		fmt.Println("client : Read from chout : ", v)
		fmt.Println("client : Content of b.slice : ", b.slice)
		b.readNext() //send msg to go routine to continue looping, and get another value from in channel.
	}
}

//buffer is a buffer
type buffer struct {
	chOut          chan int  //chout, the channel out to be read by client
	slice          []int     //slice which is the actual buffer
	confirmNewRead chan bool //used to wait for confirmation of grabbing the next value from input channel.
	size           int       //size of buffer
}

//NewBuffer create a new buffer
func NewBuffer(m int) *buffer {
	return &buffer{
		chOut:          make(chan int),
		confirmNewRead: make(chan bool),
		size:           m,
	}
}

//start will start filling the buffer, to continue filling buffer  use the readNext method.
func (b *buffer) start(chIn chan int) {
	go func() {
		for len(b.slice) < b.size-1 {
			v, ok := <-chIn
			if !ok {
				//log.Println("SERVER: done reading chIn in the slice filling loop at the top")
				break
			}
			b.slice = append(b.slice, v)
		}

		//Loop and read another value as long as the slice is > 0.
		// Since we fill the buffer when we start as the first thing
		// the only reason for the length of the slice is 0 is that
		// the input channel is closed, and the decrement of the channel
		// value by value have started.
		for len(b.slice) > 0 {
			v, ok := <-chIn

			//input channel closed ?
			if !ok {
				b.slice = b.slice[1:]
			}

			//length of slice already full ?
			if len(b.slice) == b.size {
				b.slice = b.slice[1:]
				b.slice = append(b.slice, v)
			}

			//input channel not closed, and slice buffer not filled to size.
			if ok && len(b.slice) < b.size {
				b.slice = append(b.slice, v)
			}

			//slice have been emptied, break out of for loop.
			if len(b.slice) == 0 {
				break
			}

			b.chOut <- b.slice[0]
			//Wait for confirmation for another read from main.
			<-b.confirmNewRead
		}
		close(b.chOut)
	}()
}

//readNext will, relese the lock on the go routine inside the start method,
//and let it read another value from the incomming channel and put it
//into the buffer.
func (b *buffer) readNext() {
	b.confirmNewRead <- true
}
