package main

import (
	"fmt"
)

var a, c string
var b MinNyeType

//MinNyeType just a type for testing
type MinNyeType string

//New is for creating a new minNyeType
func New(s string) MinNyeType {
	return MinNyeType(s)
}

func main() {
	a = "apekatt"
	b = "grevling"
	fmt.Printf("a = %v, og a er av typen = %T \n", a, a)
	fmt.Printf("b = %v, og b er av typen = %T \n", b, b)

	c = string(b)
	fmt.Printf("c = %v, og c er av typen = %T \n", c, c)

	d := New(a)
	fmt.Printf("d = %v, og d er av typen = %T \n", d, d)
}
