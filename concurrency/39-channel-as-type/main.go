package main

import "fmt"

type chType chan int

func (c chType) len() int {
	return len(c)
}

func main() {
	intF := func() int {
		return 1
	}

	ch := make(chType, 10)
	ch <- intF()
	fmt.Printf("Length of channel : %v\n", ch.len())
}
