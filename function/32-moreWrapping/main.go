package main

import (
	"fmt"
	"log"
)

// structure to hold the wrapped function
type funcRunner struct {
	// We want to Run any function that takes no inputs,
	// and return an error
	fønk func() error
}

// Return a new funcRunner type with a Run method to start
// the wrapped function.
func newFuncRunner(f wrapperFunc) funcRunner {
	fr := funcRunner{
		fønk: f,
	}
	return fr
}

// Run method to start the wrapped function.
func (f funcRunner) Run() error {
	err := f.fønk()
	if err != nil {
		return err
	}

	return nil
}

// Describe the signature of a wrapper func. It takes no
// inputs, and only returns an error.
type wrapperFunc func() error

// adder orchestrates and returns a function with the signature
// of a wrapperFunc.
func adder(a int, b int) wrapperFunc {
	// We define even another funtion to be wrapped inside
	// the wrapperFunc returned.
	fToWrap := func() {
		log.Printf("Executing wrapped function\n")
	}

	return func() error {
		// Add the numbers.
		fmt.Printf("adding %v+%v=%v\n", a, b, a+b)
		// Execute the other function we also wrapped in.
		fToWrap()
		return nil
	}
}

func main() {
	// Use a predefined wrapper function called adder.
	{
		fr := newFuncRunner(adder(10, 10))

		err := fr.Run()
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}

	// Create a literal wrapper function where we wrap a channel
	// into the function to feed the function a value.
	{
		var wf wrapperFunc
		// Prepare a channel to feed values into the function later.
		ch := make(chan int)
		// This wrapperFunction starts a Go routine when called,
		// which will read a value from the channel, do some, and
		// terminate.
		wf = func() error {
			go func() {
				v := <-ch
				fmt.Printf("adding %v + %v = %v\n", v, v, v+v)
			}()
			return nil
		}

		fr := newFuncRunner(wf)
		err := fr.Run()
		if err != nil {
			log.Printf("error: %v\n", err)
		}

		// Feed some value into the function we wrapped earlier.
		ch <- 10
	}

}
