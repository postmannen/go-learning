package main

import "fmt"

func anotherFunction(f func(string) string) {
	result := f("Animal")
	fmt.Println(result)
}

func main() {
	myFunc := func(s string) string {
		return "Hello, " + s
	}

	anotherFunction(myFunc)
}
