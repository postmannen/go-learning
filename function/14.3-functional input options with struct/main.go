/*
Function type, and function who returns a function.
 NB: The important thing to remember here is that it is the returned function
 functions signature that have to match the myFunc type, which is being used
 as input in functionExecutor. The called function can look whatever way it wants,
 just as long as the returned function matches to myFunc signature.

 The idea here is to test functions as input parameters in a function when beeing called.
*/

package main

import "fmt"

//myFunc is a function type with no input and no return values.
type myFunc func(*person)

type person struct {
	name string
	age  int
}

//setName takes the name of type string as input, but it returns
// a function that takes p *person as input parameter. That means
// we have to give it a person type as input where the returned
// function is executed (look below in the function NewPerson).
func setName(s string) myFunc {
	//Here we specify the input parameter that needs to be provided
	// where the returned function is beeing executed.
	return func(p *person) {
		p.name = s
	}
}

func setAge(n int) myFunc {
	return func(p *person) {
		p.age = n
	}
}

//NewPerson takes a myFunc as it input argument.
// Since mFunc is a variadic input, meaning we specify more than one when calling it,
// we loop over all the functions that might be given as input when calling NewPerson
// below with the `for range`loop to get each one, and call them.
func NewPerson(s string, mFunc ...myFunc) *person {
	p := &person{}

	//Here we finally give p as input parameter to the function thats beeing executed.
	for _, myF := range mFunc {
		myF(p)
	}

	return p
}

func main() {
	//Here we call NewPerson, and are giving it several functions to execute
	// up on creation.
	p1 := NewPerson("Even", setName("Harald"), setAge(140))

	fmt.Println("p1 now contains : ", p1)

}
