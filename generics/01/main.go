package main

import "fmt"

type testInterface interface {
	human | animal
}

func test1[T testInterface](value T) {
	fmt.Printf("%v\n", value)
}

// ---

type human struct {
	name string
}

type animal struct {
	name string
}

func main() {
	test1(human{"some"})
}
