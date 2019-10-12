package main

import "fmt"

func main() {
	a := []string{"monkey", "donkey"}
	fmt.Println("initially a contains : ", a)

	// make a new variable b based on 'a'. This variable will copy the header of the original
	// 'a', and have a pointer to the same backing array as 'a'. That means, any changes to the new variable 'b' will also change the content of 'a' since they share the same backing array.
	b := a

	// make a new variable 'c' which is the same size as 'a'.
	c := make([]string, len(a))
	// and make a copy of the content of 'a'. 'c' will not have any reference to 'a', and 'c'
	// will not share the backing array of a. That means any changes that happens to either 'a'
	// or 'b' in this code will not change the content of 'c'.
	copy(c, a)

	// now lets change 'b', which also will change 'a' since they share the same backing array.
	b[1] = "zebra"

	// as we can see, the change of b which is a reference to a, also changed the content of a.
	fmt.Println("after change of b, a contains : ", a)
	// since c is a copy of a before it got changed, and contains no reference to a, it was not
	// changed when we changed b
	fmt.Println("c which is a copy of the original a, contains : ", c)

}
