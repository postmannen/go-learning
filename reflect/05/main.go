//Testing example from https://www.youtube.com/watch?v=vq_LnLViCjY
package main

import (
	"fmt"
	"reflect"
)

//A is a struct
type A struct {
	NameField  string
	ValueField string `rename:"val"`
}

func main() {
	a := A{NameField: "a name", ValueField: "a value"}

	v := reflect.ValueOf(a)
	fmt.Printf("%#v\n", v)

	vt := v.Type()
	fmt.Printf("%v\n", vt)

	//NumFields returns the number of fields in a struct.
	fmt.Println("Number of fields in a = ", v.NumField())

	for i := 0; i < v.NumField(); i++ {
		f := vt.Field(i)
		name := f.Name //.Name will give the name of the field inside the struct
		if t, ok := f.Tag.Lookup("rename"); ok {
			name = t
		}
		fmt.Printf("%s: %s\n", name, v.Field(i))
	}
}
