package main

import "fmt"

func chFiller(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

const bufSize = 3

func main() {
	ch := make(chan int)
	go chFiller(ch) //Start the producing av data to the channel
	buf := []int{}

	readyCh := make(chan bool)
	for {
		//Read one from the inncomming channel
		n, ok := <-ch
		if !ok {
			fmt.Println("channel is nil")
			break
		}
		buf = append(buf, n)

		//Check if the buffer is full, shift all to the left, and make the last spot open
		go func() {
			if len(buf) > bufSize {
				buf = buf[1:]
			}
			readyCh <- true
		}()

		<-readyCh
		fmt.Println(buf)
	}

	fmt.Println("*length of buffer = ", len(buf))

}
