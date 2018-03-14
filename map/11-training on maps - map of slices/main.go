package main

import (
	"fmt"
)

func main() {
	map1 := map[string][]int{}
	map1["ku"] = []int{1}
	map1["hest"] = []int{10, 20, 30}

	fmt.Println("Content of map1 : \n", map1)
}
