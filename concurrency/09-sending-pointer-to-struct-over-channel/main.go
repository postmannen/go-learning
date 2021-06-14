package main

import (
	"fmt"
)

type myStruct struct {
	myField1 int
	myField2 int
}

func main() {
	ch := make(chan *myStruct, 1)

	go makeSome(ch)

	for v := range ch {
		fmt.Printf("%v,%T\n", v, v)
	}

}

func makeSome(c chan *myStruct) {

	for i := 1; i <= 10; i++ {
		c <- &myStruct{myField1: i, myField2: i}
	}
	close(c)
}
