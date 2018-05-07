package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "This is a string\n"
	fmt.Print(myString)

	//remove any new lines in the string, then print it
	trimmedString := strings.Trim(myString, "\n")
	fmt.Print(trimmedString)

}
