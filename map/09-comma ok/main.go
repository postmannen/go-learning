package main

import (
	"fmt"
)

func main() {
	a := make(map[string]int)
	a = map[string]int{
		"Elg":    200,
		"Gris":   50,
		"Hest":   180,
		"Giraff": 450,
	}

	height, ok := a["grevling"]

	fmt.Println(a)
	fmt.Println("Height = ", height, " and ok = ", ok)
}
