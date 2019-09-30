package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var i int
	fmt.Println("size of int = ", unsafe.Sizeof(i))

	var structType0 struct{}
	fmt.Println("size of struct type with 1 integer field = ", unsafe.Sizeof(structType0))

	var structType1 struct {
		nrOne int
	}
	fmt.Println("size of struct type with 1 integer field = ", unsafe.Sizeof(structType1))

	var structType2 struct {
		nrOne int
		nrTwo int
	}
	fmt.Println("size of struct type with 2 integer fields = ", unsafe.Sizeof(structType2))

	myEmptySlice := []byte{}
	fmt.Println("size of empty slice = ", unsafe.Sizeof(myEmptySlice))

}
