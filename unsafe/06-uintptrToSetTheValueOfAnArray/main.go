package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [2]int8{10, 0}
	fmt.Println(arr)

	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + 1)) = 11

	fmt.Println(arr)
}
