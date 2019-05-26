package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	name string
}

func main() {
	s := &myStruct{}
	fmt.Println(s)
	fmt.Println(unsafe.Sizeof(s))

	sp := (*myStruct)(nil)
	fmt.Println(sp)
	fmt.Println(unsafe.Sizeof(sp))
}
