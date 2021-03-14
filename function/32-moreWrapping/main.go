package main

import (
	"fmt"
	"log"
)

// structure to execute handler functions
type handlerRunner struct {
	runs int
}

// Return a new funcRunner type with a Run method to start
// a handler function.
func newHandlerRunner() *handlerRunner {
	fr := handlerRunner{}
	return &fr
}

// Run a handler takes a handler function as input, and
// returns an error indicating of the run where succesful
// or not.
func (f *handlerRunner) Run(w handler) error {
	err := w()
	if err != nil {
		return err
	}
	f.runs++

	return nil
}

// Format the output when printing the handlerRunner struct.
func (f *handlerRunner) String() string {
	return fmt.Sprintf("functions executed = %v", f.runs)
}

// Describe the signature of a wrapper func. It takes no
// inputs, and only returns an error.
type handler func() error

// adder orchestrates and returns a function with the signature
// of a handler which we can feed to the handlerRunner.Run
func adder(a int, b int) handler {
	// We can define even another funtion to be wrapped inside
	// the handler returned, that will be executed when the
	// handler function is executed.
	logFunc := func() {
		log.Printf("Executing some log function here.............\n")
	}

	// Return a function with the signature of a handler function type.
	return func() error {
		// Add the numbers.
		fmt.Printf("adding %v+%v=%v\n", a, b, a+b)
		// Add other function we also wrapped in to be executed.
		logFunc()
		return nil
	}
}

func main() {
	hr := newHandlerRunner()

	// --- Example 1

	// Use a predefined handler function called adder, to be called via
	// the handlerRunner.
	{
		err := hr.Run(adder(10, 10))
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}

	// --- Example 2

	// Use a literal handler function called adder, to be called via
	// the handlerRunner, where we wrap a channel into the function
	// to feed the function a value when it is executed.
	{
		var h handler
		// Prepare a channel to feed values into the function later.
		ch := make(chan int)
		// This wrapperFunction starts a Go routine when called,
		// which will read a value from the channel, do some, and
		// terminate.
		h = func() error {
			go func() {
				v := <-ch
				fmt.Printf("adding %v + %v = %v\n", v, v, v+v)
			}()
			return nil
		}

		err := hr.Run(h)
		if err != nil {
			log.Printf("error: %v\n", err)
		}

		// Feed some value into the function we wrapped earlier.
		ch <- 20
	}

	fmt.Println(hr)
}
