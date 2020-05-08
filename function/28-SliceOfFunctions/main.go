package main

import "fmt"

// Since a function is a first class citizen in Go, we can define
// a new type which is a slice of functions.
type sFu []func()

func main() {
	// Declare a new variable of the function slice type.
	var f sFu

	// Append a couple of functions to it.
	f = append(f, func() { fmt.Println("First One") })
	f = append(f, func() { fmt.Println("Second One") })

	// Range over the slice, and execute the functions, 1 by 1.
	for _, fu := range f {
		fu()
	}

}
