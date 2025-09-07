package main

import (
	"fmt"
	"unsafe"
)

func main() {
	f := func(n int) int {
		return n + 1
	}

	// Taking an unsafe pointer to a function.
	uf := unsafe.Pointer(&f)

	// Casting back to function done in 2 steps.
	ef := (*func(int) int)(uf)
	fmt.Printf("%v\n", (*ef)(1))

	// Casting back to function done in 1 step.
	fmt.Printf("%v\n", (*((*func(int) int)(uf)))(2))
}
