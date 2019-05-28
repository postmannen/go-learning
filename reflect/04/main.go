package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int
	x = 10
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	fmt.Println("ValueOf x = ", v, "TypeOf x = ", t)

}
