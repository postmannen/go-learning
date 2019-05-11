package main

import "fmt"

//addSome will prepare a function who adds two numbers, and returns the result.
// it will also embed the string s from the outer function, and prepare it
// for use for when the inner function is called.
func addSome() func(int, int) int {
	s := "doing some calculations"
	return func(a int, b int) int {
		fmt.Println(s)
		return a + b
	}
}

func main() {
	e := addSome()
	fmt.Println("-----------------------------------")
	fmt.Println(e(10, 10))
}
