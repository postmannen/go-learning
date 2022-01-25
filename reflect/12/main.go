package main

import (
	"fmt"
	"reflect"
)

type numbers struct {
	numberOne []string
}

func main() {
	n1 := numbers{numberOne: []string{"a", "b", "c", "d"}}

	valueOf := reflect.ValueOf(n1)
	fmt.Printf("valueOf contains: %#v, type: %T\n", valueOf, valueOf)

	for i := 0; i < valueOf.NumField(); i++ {
		fmt.Printf("valueOf.Field:%v,  reflect type: %T\n", valueOf.Field(i), valueOf.Field(i))

		refSl := valueOf.Field(i)
		for i := 0; i < refSl.Len(); i++ {
			fmt.Println(refSl.Field(i))
		}
	}
}
