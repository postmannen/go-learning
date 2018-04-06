package main

import (
	"fmt"
)

type myStruct struct {
	myString string
}

func myFuncPrint(m *myStruct) {
	fmt.Println(m.myString)
}

func myFuncCreate(m *myStruct) myStruct {
	return *m
}

func main() {
	fmt.Println("----------------------------TEST1---------------------------------------")
	//just call the function with a pointer to  a myStruct type without storing a variable of the type
	myFuncPrint(&myStruct{myString: "This is a test string"})

	fmt.Println("----------------------------TEST2---------------------------------------")
	//call a function with a pointer to a myStruct type, and return a non-pointer of the same type, and assign it to a variable
	varOfmyStruct := myFuncCreate(&myStruct{myString: "Another test string"})
	fmt.Printf("The new variable varOfmyStruct is of type = %T, and the myString field contains = %v\n", varOfmyStruct, string(varOfmyStruct.myString))

}
