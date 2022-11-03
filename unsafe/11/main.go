package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	aString := [8]byte{'a', 's', 't', 'r', 'i', 'n', 'g'}
	upStr := (uintptr)(unsafe.Pointer(&aString))

	sh := reflect.SliceHeader{Data: upStr, Len: len(aString)}

	s := *(*string)(unsafe.Pointer(&sh))

	fmt.Printf("%v, %T\n", s, s)
}
