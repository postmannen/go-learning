package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{5, 6, 1, 2, 7, 8, 3, 4, 9, 0}
	sort.Ints(n)
	fmt.Println(n)

	s := []string{"badger", "anaconda", "rhino", "duck"}
	sort.Strings(s)
	fmt.Println(s)

}
