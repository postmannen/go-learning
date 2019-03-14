/*
 The purpose of this exercise is to test how to compose a function based on
 what input you give to the function composer.
 Here we have two functions:
  - One who simulates the writing to a channel
  - One who simulates the writing to the console
 We then have a newOutput channel who will take what kind of output we want
 as input, and return a new function based on the kind that was chosen.

 The new function is then executed, and should give the correct output based on what
 output kind that was chosen.

*/
package main

import "fmt"

//writeToChannel is specifying how to write to a channel
func writeToChannel(s string) {
	fmt.Println("Writing to channel: ", s)
}

//writeToConsole is specifying how to write to a channel
func writeToConsole(s string) {
	fmt.Println("Writing to console: ", s)
}

type outputType int

const (
	toChannel outputType = iota
	toConsole
)

//newOutput takes and outputType as it's input, and will return a function that will
// produce that kind of output. The returned function can then be called in main.
func newOutput(ot outputType) func(string) {
	if ot == toChannel {
		return writeToChannel
	}

	return writeToConsole
}

func main() {
	s := "1234567890"

	// Create two functions f1 that will write to a channel,
	// and f2 that will write to the console, and call them.
	f1 := newOutput(toChannel)
	f1(s)
	f2 := newOutput(toConsole)
	f2(s)

}
