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

//Ett interface er en type, og kan igjen bruke alle typene til methodene interfacet har.
var i inter = myInt{variabel: 10}

//var i inter

func main() {

	i.Print()
}
