/*
 The purpose of this test is to convert one type of channel to a channel of
open interface, interface{}.
Since a chan int, or chan string, or chan inteface{} is it's own chan type,
we can't directly do `openInterfaceChannel <- intChannel`, we need to iterate
over each item received, and put that value on the new out channel.
To figure out what kind of channel we have as input we use the reflect.ValueOf, and use
that to set the type of the incomming channel, and then put that value of a new
out channel which is of type interface{}.
*/
package main

import (
	"fmt"
	"log"
	"reflect"
)

//incommingSignal will be the remote system that sends us some values
// of a specific type in this test.
func incommingSignal() chan int {
	ch := make(chan int)
	go func() {
		for i := 200; i <= 210; i++ {
			t := i
			fmt.Printf("Putting original value=%v, of type %T on channel\n", t, t)
			ch <- t
		}
		close(ch)
	}()

	return ch
}

//convToOpenInterface will take any value as input, and convert it to a `chan interface{}`
// It checks if it is a channel, if it is a channel we take the current value
// of the in channel of type int and put it on the out channel which is of type interface{}
func convToOpenInterface(ci interface{}) chan interface{} {
	cin := reflect.ValueOf(ci)
	fmt.Printf("cin = %#v\n", cin) //cin = (chan int)(0xc000086000)
	cout := make(chan interface{})

	go func() {
		for {
			x, ok := cin.Recv()
			if !ok {
				log.Printf("--- convToOpenInterface: not a channel %v, closing cout\n", ok)
				close(cout)
				break
			}
			fmt.Printf("* Converting to x %v %T, and putting on cout channel\n", x, x)

			cout <- x
		}
	}()
	return cout
}

func main() {
	inChannel := incommingSignal()
	convInChannel := convToOpenInterface(inChannel)
	doSome(convInChannel)

}

//doSome will loop over the values of the new out channel created, and we can now
// see that they are no longer of type int but of type interface{}
func doSome(ch chan interface{}) {
	for v := range ch {
		fmt.Printf("incomming of type %T, contains %v\n", v, v)
	}
}
