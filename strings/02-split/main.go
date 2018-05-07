package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "one two three"

	myNewSLice := strings.Split(myString, " ")
	for i, v := range myNewSLice {
		fmt.Println(i, v)
	}

}
