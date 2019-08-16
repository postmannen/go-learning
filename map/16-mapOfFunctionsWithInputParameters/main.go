/*
The purposed of this example is to test out creating a map of functions,
where all the function input parameters are pre declared, so you can
do fn() directly on the map value instead of having to fill it with input
values after the map value have been received.
*/

package main

import (
	"fmt"
)

// chopAnimals are a function that takes a string, and returns the modified string,
// and this is the actual function we want to have as our map value..filled with
// it's correct input parameter/values.
func chopAnimal(s string) string {
	r := fmt.Sprintf("The %v have been chop'ed !\n", s)
	return r
}

func main() {
	// What we care about with our function map values are the return values.
	// That is why we initialize the map with 'map[string]func()string', and
	// not 'map[string]func(string)string' (this took some time to understand :-).
	// Then when we assign a value to the map value we create an unnamed function
	// which returns the string which is the same as our original function does.
	// Then within that function body we call the original function we want to
	// call with it's input string value set, and the returned string are returned
	// all the way out.
	chopMap := make(map[string]func() string)
	chopMap["monkey"] = func() string { return chopAnimal("monkey") }

	// get the map value for the key "monkey". The returned value should be a
	// function with the signature of func()string, and our original/wanted function
	// we want to call is wrapped inside that function.
	s := chopMap["monkey"]
	// We can then call our function without any input parameters, which again
	// will call the inner wrapped function, and the return value which is a string
	// will be returned out.
	fmt.Println(s())
}
