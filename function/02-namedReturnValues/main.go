package main

import "fmt"

func main() {
	fmt.Println(calulate(100, 200))
	sum1, sum2 := calulate(1000, 2000)
	fmt.Println(sum1, sum2)
}

//calculate adds and multiplicates two variables, returns the add sum and multiplication sum
func calulate(x int, y int) (sumAdded int, sumMultiplicated int) {
	add := x + y
	mul := x * y
	return add, mul
}
