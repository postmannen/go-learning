package main

import "fmt"

// number will hold the number being calculated on.
type number struct {
	sum int
}

// add will take a number to add, and return a function with the
// signature func(*number).
// It will create a new function to return which takes a 'n *number'
// as it's input, and add's n1 to that number. The argument n is the
// variable of type number we add to the function when calling it
// within the loop of the calculator function.
func add(n1 int) func(*number) {
	return func(n *number) {
		n.sum += n1
	}
}

// sub will take a number to subtract, and return a function with the
// signature func(*number).
func sub(n1 int) func(*number) {
	return func(n *number) {
		n.sum -= n1
	}
}

// calculator takes the initial value for the calculation, and and
// arbitrary number of sub or add functions to execute in order on
// the initial number, and in the end it will print the sum of the
// calculations.
// Since the input functions is a variadic type we can range over
// that type, get a single function, and execute it on the sum.
func calculator(n int, fn ...func(*number)) {
	num := number{
		sum: n,
	}

	for _, f := range fn {
		f(&num)
	}

	fmt.Println("sum = ", num.sum)
}

func main() {
	// Since the calculator function takes a variadic value of the
	// function type as input, we can add as many add or sub
	// functions as we want.
	calculator(130, add(100), sub(500))

}
