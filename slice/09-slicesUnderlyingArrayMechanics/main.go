package main

import (
	"fmt"
)

func main() {
	// Create a slice of string, length 0, capacity 20, and fill it with some values
	a := make([]string, 0, 20)
	a = append(a, "a", "b", "c", "d", "e", "f")
	fmt.Printf("len a = %v, cap a = %v, a = %v, &a[0] = %v\n", len(a), cap(a), a, &a[0])

	// Slice out the first 4 values, and notice that the address of index 0, is the
	// same as the original slice a.
	b := a[0:4]
	b = append(b, "x")
	fmt.Printf("len b = %v, cap b = %v, b = %v, &b[0] = %v\n", len(b), cap(b), b, &b[0])
	fmt.Printf("len a = %v, cap a = %v, a = %v, &a[0] = %v\n", len(a), cap(a), a, &a[0])
	fmt.Printf("\n")

	// Slice out the first 4 values, but with the capacity set to only what is being sliced.
	// When we assign this to c we will have created a slice sharing the same backing
	// array as the original slice, but only with a capacity == length of slice.
	// Then when we append to c we will need to extend the capacity of the slice to hold the
	// extra value, and a new backing array will be created, and the old values will be copied
	// over. As we can see it no longer shares the original backing array.
	c := a[0:4:4]
	c = append(c, "z")
	fmt.Printf("len c = %v, cap c = %v, c = %v, &c[0] = %v\n", len(c), cap(c), c, &c[0])
	fmt.Printf("len a = %v, cap a = %v, a = %v, &a[0] = %v\n", len(a), cap(a), a, &a[0])
}
