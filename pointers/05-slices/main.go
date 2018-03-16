package main

/*
Testing altering slices via methods
*/

import (
	"fmt"
)

type mySlice []int

//This one does not alter the data since using a range loop makes a copy of the value, and puts it into v
func (m mySlice) addOne() {
	for _, v := range m {
		v++
	}
}

//This one works, and alters the data since it is directly accessing the value based on the index of the slice
func (m mySlice) addOneMore() {
	fmt.Println("The len of m = ", len(m))
	for i := 0; i < len(m); i++ {
		m[i]++
	}
}

func main() {
	aSlice := mySlice{10, 20, 30}
	fmt.Println("The initial slice contains = ", aSlice)
	aSlice.addOne()

	fmt.Println(aSlice)

	fmt.Println("---------NEXT TEST--------")
	aSlice.addOneMore()
	fmt.Println(aSlice)
}
