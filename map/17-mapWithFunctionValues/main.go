package main

import (
	"fmt"
)

func returnInt(i int) int {
	return i
}

func main() {
	// Create a map with functions to call as values.
	// To actually call the function add parantheses after defining the
	// map key, and specify the value for the function
	m1 := map[string]func(int) int{
		"ape": returnInt,
	}

	//Here we call the m1 function value and give the function an int value
	// as input.
	fmt.Println(m1["ape"](10))

	// If we want to predefine some values for the function we need to
	// change the function signature for the map. The reason is that we
	// no longer have a function which takes an int as input and returns
	// and int, we only care about a function who returns an int valie.
	// Since we're filling in the input values for the function directly
	// we leave out the input field for the function in the map declaration,
	// and instead returning a function with the value wanted given as input.
	m2 := map[string]func() int{
		"ape": func() int {
			return returnInt(10)
		},
	}

	// Here we call the m2 function value, but since the input is already
	// predefined we don't give an input value when calling the function value.
	fmt.Println(m2["ape"]())

}
