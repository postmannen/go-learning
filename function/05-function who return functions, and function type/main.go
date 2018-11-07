package main

import (
	"fmt"
)

// --------------------------------------------------------------
//First we define a func who takes no arguments,
// and return another func.
func aFunc() func() int {
	//That other func takes no arguments, and returns an INT,
	// just like we specified above
	return func() int {
		return 47
	}
}

// --------------------------------------------------------------
type bAdd func() int

func main() {
	//Here we assign the function "aFunc" to "a", and execute a.
	a1 := aFunc()
	fmt.Printf("*a* - Return from a1 which now is aFunc : %v, type %T \n", a1(), a1)
	fmt.Printf("*a* - a1 is of type : %T \n", a1)

	//we can also assign "a" to another variable, and exexute that variable.
	a2 := a1
	fmt.Printf("*a* - Return from a2 which now is aFunc : %v, type %T \n", a2(), a2)
	fmt.Printf("*a* - a2 is of type : %T \n", a2)

	// -----------------------------------------------------------
	fmt.Println("----------------------------------------------------------------")

	//Since bAdd and aFunc share the same signature, we can use aFunc with bAdd
	//b1 will declared as type bAdd
	var b1 bAdd = aFunc()
	fmt.Printf("*b* - Return from b1 which now is bAdd : %v, type %T \n", a2(), a2)
	fmt.Printf("*b* - b1 is of type : %T \n", b1)

	// -----------------------------------------------------------
	fmt.Println("----------------------------------------------------------------")

	//create a function that returns a new function of the type bAdd
	c1 := func() bAdd {
		fmt.Println("Executing the outer function")

		//This inner function has the same signature as bAdd
		return func() int {
			fmt.Println("Executing the inner function")
			return 100
		}
	}

	//c1 should now contain a pointer to the inner function above which is of type bAdd
	//We can also see on the output that the inner function is not executed.
	fmt.Printf("c1 - c1 is of type : %T, and the value returned is %v\n", c1, c1())
	//The value returned seems like a pointer, and probably is a pointer
	// to the inner func. Lets check
	c2 := c1()
	fmt.Printf("c2 - c2 is of type : %T, and the value returned is %v\n", c2, c2())
	//On the out put we can now see that both the inner and outer functions
	// are executed.

	// -----------------------------------------------------------
	fmt.Println("----------------------------------------------------------------")
}
