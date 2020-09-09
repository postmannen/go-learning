package main

import (
	"fmt"
	"reflect"
)

func insertA(i int) {
	var anInt int
	anInt = i
	_ = fmt.Sprintln(anInt)

}

func insertB(i interface{}) {
	v := reflect.ValueOf(i).Interface().(int)
	var anInt int
	anInt = v
	_ = fmt.Sprintln(anInt)

	//var anInt int
	//anInt = i
	//_ = anInt
}

func main() {
	insertB(int(10))
}
