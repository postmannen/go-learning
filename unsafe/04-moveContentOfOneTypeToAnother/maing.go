package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a uint64
	a = 1000

	var b int64

	//How this works:
	//
	// If two values of any type share the same memory layout, they can be converted with
	// the unsafe package. uint64 and int64 are both 64bits and thus share the same
	// memory layout.
	//
	// Any pointer to value can be converted to unsafe pointer, hence : unsafe.Pointer(&a)
	// Any unsafe pointer can be converted to pointer to any type, hence: (*int64)(unsafe.Pointer(&a))
	// Then we dereference the last pointer value, hence : b = *(*int64)(unsafe.Pointer(&a))
	//
	// &f takes a pointer to the float64 value stored in f.
	// unsafe.Pointer(&f) converts the *float64 to an unsafe.Pointer.
	// (*uint64)(unsafe.Pointer(&f)) converts the unsafe.Pointer to *uint64.
	// *(*uint64)(unsafe.Pointer(&f)) dereferences the *uint64, yielding a uint64 value.
	b = *(*int64)(unsafe.Pointer(&a))

	fmt.Println(b)

}
