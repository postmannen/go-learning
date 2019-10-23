package main

import "fmt"

func main() {
	x := []string{"Donald", "Duck", "Andeby", "Norge"}
	y := []string{"Fantomet", "Walker", "Hodeskallegrotten", "Afrika"}
	fmt.Println(x)
	fmt.Println(y)

	// x[:1] will keep just the first field,
	// x[2:] will keep everything after slice index 2,
	// The result will be that you throw away the first item of the slice.
	x = append(x[:1], x[2:]...)

	fmt.Println("Removing index 0 = ", x)

	// Remember ... if you want to append one slice to another.
	x = append(x, y...)
	fmt.Println("When y is appended to x, x becomes : ", x)

	fmt.Println("´n")

	// Range is the same as foreach in Perl.
	for i, v := range x {
		fmt.Printf("%v\t%v\n", i, v)
	}

}
