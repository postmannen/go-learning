package main

import "fmt"

func anotherFunction(f func(string) string) {
	result := f("Animal")
	fmt.Println(result)
}

func main() {
	// Create a function my declaring av variable using an
	// anonymous function declaration.
	// This is the same as declaring a function by typing :
	// func myFunc(s string) string {.......}
	//
	// This function takes a string as an argument, and it
	// takes that string and adds "Hello" before it.
	// This function will be executed only inside the function
	// with the name anotherFunc.
	myFunc := func(s string) string {
		return "Hello, " + s
	}

	// Here we call anotherFunction, and gives our function as
	// an input. That input'ed function will then be called
	// inside anotherFunction.
	anotherFunction(myFunc)
}
