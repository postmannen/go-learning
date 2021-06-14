/*
 A litte practise with maps to count the occurence of numbers.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numMap := map[int]int{}

	for i := 1; i < 100; i++ {
		n := rand.Intn(10)

		_, keyExist := numMap[n]
		if !keyExist {
			numMap[n] = 1
		} else {
			numMap[n]++
		}

	}

	for k, v := range numMap {
		fmt.Println(k, v)
	}
}
