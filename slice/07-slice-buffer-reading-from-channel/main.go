package main

import (
	"fmt"
)

func main() {
	chIn := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			chIn <- i
		}
		close(chIn)
	}()

	b := NewBuffer()
	b.start(chIn)

	for v := range b.chOut {
		fmt.Println("client : Read from chout : ", v)
		fmt.Println("client : Content of b.slice : ", b.slice)
		b.readNext() //send msg to go routine to continue looping, and get another value from in channel.
	}
}

//buffer is a buffer
type buffer struct {
	chOut          chan int
	slice          []int
	confirmNewRead chan bool
}

//NewBuffer create a new buffer
func NewBuffer() *buffer {
	return &buffer{
		chOut:          make(chan int),
		confirmNewRead: make(chan bool),
	}
}

//start will start filling the buffer, to continue filling buffer  use the readNext method.
func (b *buffer) start(chIn chan int) {
	var max = 3
	go func() {
		for len(b.slice) < max-1 {
			v, ok := <-chIn
			if !ok {
				//log.Println("SERVER: done reading chIn in the slice filling loop at the top")
				break
			}
			b.slice = append(b.slice, v)
		}
		//fmt.Println("SERVER: done with the first fill, slice looks like , ", b.slice)
		//fmt.Println("-------------------------------------------------------------------")

		for len(b.slice) > 0 {
			v, ok := <-chIn
			//fmt.Println("SERVER: just read chIn = ", v)

			if !ok {
				//fmt.Println("SERVER: IF 1, done reading chIn,removing first item from slice")
				b.slice = b.slice[1:]
			}

			if len(b.slice) == 3 {
				//fmt.Println("SERVER: IF 2, removing first of slice, adding to the end of slice, ", v)
				b.slice = b.slice[1:]
				b.slice = append(b.slice, v)
			}

			if ok && len(b.slice) < max {
				//fmt.Println("SERVER: IF 3, appending to slice, ", v)
				b.slice = append(b.slice, v)
			}

			if len(b.slice) == 0 {
				//fmt.Println("SERVER: IF 4, length of slice = 0, breaking out of for loop")
				break
			}
			//fmt.Printf("SERVER: END OF FOR: before sending %v to client, slice contains = %v\n", b.slice[0], b.slice)
			b.chOut <- b.slice[0]
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
