package main

import "fmt"

// add is a function that in the end will add 2 numbers.
// It will take the value x as input when called, and return a new
// function to be called which takes the second value y that will
// be added to x.
//
// The proper use of this function is that you initialize it to a
// variable and give it the value for x. That variable is now an
// add function containing the value of x.
// Since that variable now is a function of the type of the returned
// function, and looks like this 'func(y int) int' we call that
// function-variable and give it the value for y, and it will return
// the result of x and y added together.
func add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	// initialize a variable to hold the returned function by calling
	// the add function and give it the value for x.
	addSome := add(10)
	// call that returned function, and give it the value for y.
	fmt.Println(addSome(5))
}
