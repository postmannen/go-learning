package main

import (
	"fmt"
)

type customer struct {
	name    string
	surname string
}

func main() {
	// Create a slice of type customer
	p := []customer{}

	// Append values of type customer to the slice p.
	p = append(p, customer{"bob", "bobson"})
	p = append(p, customer{"arne", "arneson"})
	p = append(p, customer{"knut", "knutson"})
	fmt.Println(p)
	fmt.Println(p[1])
	fmt.Println(p[1].name, p[1].surname)

	lengthOfSlice := len(p)
	fmt.Println("Length og slice", lengthOfSlice, "\n-----------------------------")

	for i := 0; i < lengthOfSlice; i++ {
		fmt.Println(p[i].name, p[i].surname)
	}
}
