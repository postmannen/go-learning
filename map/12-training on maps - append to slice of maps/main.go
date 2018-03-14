package main

import (
	"fmt"
)

func main() {
	//test1
	slice1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	map1 := make(map[int]int)

	for i, v := range slice1 {
		fmt.Println("Content of slice1 = ", v)
		map1[i] = v
	}

	fmt.Println("map1 = ", map1)

	//test2
	map2 := map[string][]int{}

	//assign the complete slice1 to the key 'numbers'
	map2["numbers"] = slice1

	//append a number to the slice which is a value of the key 'numbers'
	map2["numbers"] = append(map2["numbers"], 100)

	fmt.Println("------------------")
	fmt.Println("map2 contains = \n", map2)

}
