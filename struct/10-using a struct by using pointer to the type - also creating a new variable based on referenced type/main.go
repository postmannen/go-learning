package main

import (
	"fmt"
)

type myStruct struct {
	myString string
}

//myFuncPrint, just prints out the content of myStructs field myString
func myFuncPrint(m *myStruct) {
	//...do something with the struct data given as input
	fmt.Println(m.myString)
}

//myFuncPrint, takes a pointer to struct as input, and returns a non-pointer to the same struct type
func myFuncCreate(m *myStruct) myStruct {
	//here we dereference the struct, so we're returning a copy of the input struct with all its present values in the struct fields.
	return *m
}

func main() {
	fmt.Println("----------------------------TEST1---------------------------------------")
	//call the function with a pointer to  a myStruct type without storing a variable of the type,
	//just directly referencing the struct type with '&', and filling a value for the field of the struct
	myFuncPrint(&myStruct{myString: "This is a test string"})

	fmt.Println("----------------------------TEST2---------------------------------------")
	//call a function with a pointer to a myStruct type, and return a non-pointer of the same type, and assign it to a variable
	varOfmyStruct := myFuncCreate(&myStruct{myString: "Another test string"})
	fmt.Printf("The new variable varOfmyStruct is of type = %T, and the myString field contains = %v\n", varOfmyStruct, string(varOfmyStruct.myString))

}
