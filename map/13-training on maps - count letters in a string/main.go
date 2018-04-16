package main

import (
	"fmt"
)

func main() {
	myString := "This is a test with some letters in it. The idea is to count the use of all the letters!"

	letterCount := make(map[int32]int)

	//range all the possible letters in the ascii table
	for i := 14; i <= 122; i++ {
		//range all the letters in the string variable
		//since a string is a []byte we can use range to get each individual letter

		for _, v := range myString {
			if i == int(v) {
				_, ok := letterCount[v]
				if !ok {
					letterCount[v] = 1
				} else {
					letterCount[v]++
				}
			}
		}
	}

	for i, v := range letterCount {
		fmt.Printf("The letter %v found %v times.\n", string(i), v)
	}
}
