package main

import (
	"fmt"
)

type myString struct {
	variabel string
}

func (v *myString) Print() {
	fmt.Println(v)
}

type myInt struct {
	variabel int
}

func (v myInt) Print() {
	fmt.Println(v)
}

type inter interface {
	Print()
}

var i inter = myInt{variabel: 10}

func main() {

	i.Print()
}
