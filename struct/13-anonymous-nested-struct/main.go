package main

import (
	"fmt"
)

func main() {
	x := struct {
		y struct {
			v int
		}
		z string
	}{
		z: "z",
		y: struct {
			v int
		}{
			v: 1,
		},
	}
	fmt.Println(x)
}
