// https://golang.org/pkg/unsafe/#Pointer
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [2]int8{10, 0}
	fmt.Println(arr)

	// By using the uintptr we are allowed to do pointer aritmetrics
	// like adding 1 to the address like we do in the example below
	// to get to the second field of the array.
	//
	// 1. create an unsafe.Pointer of the address of the first field of the array.
	// 2. Cast that unsafe.Pointer into an uintptr, and add 1 for the next field.
	// 3. Create an unsafe pointer of the new address
	// 4. Convert that unsafe pointer into a pointer to an int8 (*int8)
	// 5. Dereference that *int by *(*int8)
	// 6. We hold the array slow, and we can assign it the value 11
	//
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + 1)) = 11

	fmt.Println(arr)
}
