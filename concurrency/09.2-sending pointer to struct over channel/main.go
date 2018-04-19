/*
Test how to pass pointers to array via a channel, change it, and check if the change
is reflected in the Go routine where the array was created.

Create an array with values within one Go routine, pass a pointer to that value on a channel.
Pick it up from the channel in another Go routine, change it.
Then print them all out in the first go routine to see if it was changed.
*/
package main

import (
	"fmt"
	"sync"
)

type myStruct struct {
	myField1 int
	myField2 int
}

var wg sync.WaitGroup

const lengthArray = 10

func main() {
	ch := make(chan *myStruct, 1)

	wg.Add(1)
	go makeSome(ch)

	wg.Add(1)
	go changeSome(ch)
	wg.Wait()

}

func makeSome(c chan *myStruct) {
	myArray := [lengthArray]myStruct{}

	for i := 0; i < lengthArray; i++ {
		myArray[i].myField1 = i
		myArray[i].myField2 = i
		c <- &myArray[i]
	}
	close(c)
	fmt.Println(myArray)
	wg.Done()
}

func changeSome(c chan *myStruct) {
	for v := range c {
		v.myField1 = 2
		v.myField2 = 2
	}
	wg.Done()
}
