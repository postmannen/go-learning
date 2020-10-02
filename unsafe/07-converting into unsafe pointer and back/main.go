package main

import (
	"fmt"
	"unsafe"
)

type numbers [8]int

func main() {
	s1 := numbers{1, 2, 3, 4, 5, 6, 7, 8}

	// 1. ----------- Converting into a [8]byte which is big enough to hold the data ------------
	fmt.Printf("s1 contains = %v, and is of type %T\n", s1, s1)
	fmt.Println("------------------------------------------------------")

	{
		// Any Pointer type can be converted into an unsafe.Pointer
		//
		// converting s1 into an unsafe pointer gives
		fmt.Printf("us contains : %v, and is of type : %T\n", unsafe.Pointer(&s1), unsafe.Pointer(&s1))

		// Converting the s1 into an unsafe.Pointer, and then convert it into a *[16]byte.
		//
		// The parentheses around (*[8]byte) are needed since we are converting into a pointer
		// value. If the (...) where left out it would be interpreted as a multiplier and not
		// a pointer.
		us := (*[8]byte)(unsafe.Pointer(&s1))
		fmt.Printf("us contains = %v, and is of type %T\n", us, us)
		fmt.Printf("*us contains %v, and is of type %T\n", *us, *us)

		// Convert it back to the original struct
		s1back := (*numbers)(unsafe.Pointer(us))
		fmt.Printf("s1Back contains = %v, and is of type %T\n", s1back, s1back)
		fmt.Printf("*s1Back contains %v, and is of type %T\n", *s1back, *s1back)
	}

}
