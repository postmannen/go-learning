package main

import "fmt"

type a struct {
	b struct {
		c string
	}
}

func main() {
	var a a
	a.b.c = "ape"

	fmt.Println(a)
}
