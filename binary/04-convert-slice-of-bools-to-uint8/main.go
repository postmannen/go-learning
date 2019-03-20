package main

import (
	"fmt"
)

func main() {
	b := []bool{true, false, false, false, true, false, false, true}
	fmt.Println("The values of b which is a []bool, ", b)

	n := convert(b)

	fmt.Printf("Result, binary=%b, numeric=%v\n", n, n)
}

func convert(b []bool) uint8 {
	var n uint8
	for i := 0; i < len(b)-1; i++ {
		//fmt.Printf("loop, iterating i=%v, value=%v, type=%T\n", i, v, v)

		//If true we swap the least significant bit to 1,
		// if false we leave it to its default which is 0.
		if b[i] == true {
			n = n | 1
		}

		//Bit shift all 1 to the left, which means..we add a zero to the end
		// to make room for the next value.
		n = n << 1
		//fmt.Printf("Value is now, %b\n", n)
	}

	//instead of checking for the end after each iteration to handle the last bit,
	// we do it after, and go from 14.2 ns/op to 12.0 ns/op
	if b[len(b)-1] == true {
		n = n | 1
	}

	return n
}
