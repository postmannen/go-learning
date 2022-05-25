package main

import "fmt"

type values[T int | string] []T

func main() {
	v := values[int]{1, 2}
	fmt.Printf("%v\n", v)
}
