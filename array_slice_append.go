package main

import (
	"fmt"
)

type myStruct struct {
	first  string
	second string
}

func main() {
	a := myStruct{"hei", "hallo"}

	b := []int{1, 2, 3}
	b = append(b, 4)

	fmt.Println("Hello, playground", a, b, b[0])

}
