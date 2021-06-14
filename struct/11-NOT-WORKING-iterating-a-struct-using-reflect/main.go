package main

import (
	"fmt"
	"reflect"
)

type person struct {
	FName string
	LName string
	Place string
}

func main() {
	p1 := person{
		FName: "ole",
		LName: "Brum",
		Place: "brumplassen",
	}
	p2 := person{
		FName: "ole",
		LName: "Brum",
		Place: "brumplassen",
	}

	v1 := reflect.ValueOf(p1)
	v2 := reflect.ValueOf(p2)
	fmt.Println(v1, v2)

	for i := 0; i < v1.NumField(); i++ {
		if v1.Field(i) == v2.Field(i).Interface() {
			fmt.Printf("%v and %v type = %T, %T where equal\n", v1.Field(i), v1.Field(i), v2.Field(i), v2.Field(i))
		} else {
			fmt.Printf("%v and %v type = %T, %T where not equal\n", v1.Field(i), v1.Field(i), v2.Field(i), v2.Field(i))
		}
	}

}
