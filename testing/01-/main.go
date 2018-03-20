package main

import (
	"fmt"
)

func addByTwo(x int) (result int) {
	result = x + 2

	return
}

func main() {
	fmt.Println(addByTwo(2))
}
