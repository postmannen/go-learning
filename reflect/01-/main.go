package main

import (
	"fmt"
	"reflect"
)

var myString = "Dette er noe tekst"

type mystruct struct {
	var1 string
	var2 int
}

func main() {
	fmt.Println(reflect.TypeOf(myString))

	//myVar := mystruct{var1: "apekatt", var2: 1000}
}
