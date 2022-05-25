package main

import "fmt"

func createSlice[T any](v ...T) (result []T) {
	for _, vv := range v {
		result = append(result, vv)
	}

	return result
}

func main() {
	fmt.Printf("%#v\n", createSlice(1, 2, 3, 4))
	fmt.Printf("%#v\n", createSlice("1", "2", "3", "4"))
}
