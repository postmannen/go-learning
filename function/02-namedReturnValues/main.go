package main

import "fmt"

func main() {
	fmt.Println(calulate(100, 200))
	sum1, sum2 := calulate(1000, 2000)
	fmt.Println(sum1, sum2)
}

// calculate adds and multiplicates two variables, returns the add sum and multiplication sum
// We can add names to the return variables and make it document how the function works in a
// better way. We don't even have to use the variables the function declaration creates,
// we can just return our own ones used in the function body as long as they are of the same
// type.
func calulate(x int, y int) (sumAdded int, sumMultiplicated int) {
	add := x + y
	mul := x * y
	return add, mul
}
