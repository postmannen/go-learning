package main

import (
	"fmt"
	"strconv"
)

type Int int

func (i Int) String() string {
	iString := "int converted to string = " + strconv.Itoa(int(i))
	return iString
}

func main() {
	var n Int
	n = 100
	fmt.Println(n)
	fmt.Printf("%v\n", n)
}
