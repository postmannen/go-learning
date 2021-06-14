package main

import (
	"fmt"
	"reflect"
)

func main() {
	myString := "Apekatt"
	fmt.Println(reflect.TypeOf(myString))

	myMap := map[int]bool{}
	myReflect := reflect.TypeOf(myMap)
	fmt.Printf("%v, of type = %T \n", myReflect, myReflect)
	fmt.Println(myReflect.Elem())
	fmt.Println(myReflect.Key())
}
