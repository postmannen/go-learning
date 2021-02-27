package main

import (
	"fmt"
	"reflect"
)

type t int

func (t *t) String() string { return "hi" }
func (t *t) Unused() string { return "oops" }

func examineStringer(s fmt.Stringer) {
	// The pointer causes the resulting interface{} to record
	// the actual interface type, and then the reflect.Value
	// retains that information even when you dereference it.
	v := reflect.ValueOf(&s).Elem()
	fmt.Printf("v: type %s\n", v.Type().String())
	fmt.Printf("v: %d methods\n", v.Type().NumMethod())
}

func main() {
	var x t
	examineStringer(&x)
}
