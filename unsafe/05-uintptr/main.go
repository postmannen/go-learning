package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// The code below works, but gives a warning that it is not the correct way to do it.
	//
	//a := int(10)
	//fmt.Printf("The address of a = %v\n", &a)
	//
	//aUintptr := (uintptr)(unsafe.Pointer(&a))
	//fmt.Printf("The value of aUintptr = %x, and the type is %T \n", aUintptr, aUintptr)
	//
	//b := unsafe.Pointer(aUintptr)
	//fmt.Printf("Thr address of b = %v\n", &b)

	// The correct way to do it is to do it in one operation,
	// and don't create an uintptr variable first :
	a := int8(10)
	fmt.Printf("The address of a = %v\n", &a)

	b := (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 8))

	fmt.Printf("The address of b = %v, and the value = %v\n", &b, b)
}
