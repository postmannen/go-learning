package main

import "fmt"

type myStruct struct {
	b []int
}

func main() {
	c := myStruct{[]int{1, 2, 3}}
	c.b = append(c.b, 4, 5, 6, 7, 8, 9) //append to the slice inside the struct, struct = append([]type, ...type)

	fmt.Println("Dette er en test", c)

	for i := range c.b {
		fmt.Println(c.b[i])
	}

}
