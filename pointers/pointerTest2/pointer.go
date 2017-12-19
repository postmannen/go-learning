package main

import "fmt"

type struct1 struct {
	first1 string
	last1  string
}

type struct2 struct {
	some *struct1
}

var a struct2

func main() {

	fmt.Println("Dette er en test")

}
