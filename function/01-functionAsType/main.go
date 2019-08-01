package main

import (
	"fmt"
)

// Describe what our func type should look like.
type delFunc func()

// db is an empty struct to hold some methods.
type db struct{}

// deleteLine is a function that shares the same specifications
// as a delFunc, which is a function with no input and no return.
func (d db) deleteLine() {
	fmt.Println("Deleting line")
}

func main() {
	// db is a type that also has a method with the signature T.func(),
	// which means a type with a method with no inputs and no returns.
	var myDB db

	// Since delFun and myDB.deleteLine share the same signature we
	// should be able to assign a myDB.deleteLine to a delFunc.
	// When assigning a function to another variable as variable we
	// ommit the () at the end. If we had also written the () at the
	// end the function would have been executed, and it would have
	// been the return value we had tried to assign into a, and not
	// the function itself.
	var a delFunc = myDB.deleteLine

	// Execute the newly created function variable by adding () after
	// the variable name.
	a()

	// Check what type the variable 'a' is.
	fmt.Printf("a is of type %T\n", a)

}
