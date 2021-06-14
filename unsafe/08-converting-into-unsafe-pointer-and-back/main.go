package main

import (
	"fmt"
	"unsafe"
)

type t1 struct {
	name string
	nr   int
}

type t2 struct {
	animal string
	age    int
}

func main() {
	s1 := t1{name: "ole", nr: 7}

	s1Usp := unsafe.Pointer(&s1)
	// sine both t1 and t2 share the same structure we are able to convert either
	// to an unsafe.Pointer and then convert it over into the other.
	s2 := *(*t2)(s1Usp)

	fmt.Printf("%#v\n", s2)

}
