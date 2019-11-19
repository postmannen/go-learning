package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.Unquote("Hi"))       // Error: invalid syntax
	fmt.Println(strconv.Unquote(`Hi`))       // Error: invalid syntax
	fmt.Println(strconv.Unquote(`"Hi"`))     // Prints "Hi"
	fmt.Println(strconv.Unquote(`"Hi\x21"`)) // Prints "Hi!"

	// This will print 2 lines:
	fmt.Println(strconv.Unquote(`"First line\nSecondline"`))
}
