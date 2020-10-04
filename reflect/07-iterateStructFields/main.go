package main

import (
	"fmt"
	"reflect"
)

type numbers struct {
	numberOne int
	numberTwo float64
}

func main() {
	n1 := numbers{numberOne: 99, numberTwo: 3.14}

	n1RfType := reflect.TypeOf(n1)
	fmt.Printf("n1RfType : %v\n", n1RfType)

	fmt.Printf("numfields: %v\n", n1RfType.NumField())

	// reflect.ValueOf will returns a new value of the struct with type set to reflect.Value,
	// but will also contain the original type.
	// valueOf contains: main.numbers{numberOne:99, numberTwo:3.14}, type: reflect.Value
	valueOf := reflect.ValueOf(n1)
	fmt.Printf("valueOf contains: %#v, type: %T\n", valueOf, valueOf)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		fmt.Printf("--------------Field nr %v-----------------\n", i)
		fmt.Printf("valueOf.Field:%v,  reflect type: %T\n", valueOf.Field(i), valueOf.Field(i))
		fmt.Printf("valueOf.Type().Field(i).Type underlying type: %v\n", valueOf.Type().Field(i).Type)
	}
}
