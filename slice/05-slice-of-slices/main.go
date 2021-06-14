package main

import (
	"fmt"
)

var slice1 [][]string

func main() {
	slice1 := append(slice1, []string{"aa", "bb"})
	slice1 = append(slice1, []string{"cc", "dd"})
	fmt.Println(slice1)
}
