package main

import "fmt"

func main() {
	//If a variable is declared as an empty interface type, it can hold any type.
	// The type of 'a' will be adopted from the concrete type it holds.
	//
	var a interface{}
	a = 21
	fmt.Printf("%v, %T\n", a, a) //21, int
	a = "21"
	fmt.Printf("%v, %T\n", a, a) //21, string
}
