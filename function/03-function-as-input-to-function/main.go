package main

import (
	"fmt"
)

func addToString(s string) string {
	return fmt.Sprint(s, " apekatt")

}

// Let the function accept another function as a parameter
// As one of the parameters we accept a function with the
// signature 'func(string) string', which means we can use any
// function that takes a string as input, and returns a string.
func stringMe(name string, addS func(string) string) string {
	// The function will be executed when returning from this
	// function.
	return addS(name)
}

func main() {
	fmt.Println(stringMe("Anna", addToString))
}
