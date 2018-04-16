package main

import (
	"fmt"
)

func main() {
	var i interface{}
	i = true
	fmt.Printf("%T,%v\n", i, i)

	fmt.Println("-----------------------------------TEST1---------------------------------------")
	//to test if an empty interface holds a value, put the expected underlying value within ()
	//will set ok to true if the underlying value is of same type, and put the value into v.
	v, ok := i.(string)
	if !ok {
		fmt.Println("The value of the interface was not a string")
	} else {
		fmt.Println("The value of the interface was string,", v)
	}

	fmt.Println("-----------------------------------TEST2---------------------------------------")
	//switch can also be used with the 'type' keyword.
	switch i.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	default:
		fmt.Println("unable to detect type")
	}

	fmt.Println("-----------------------------------TEST3---------------------------------------")
	//with switch you can also get the value if the switch evaluates to true
	switch v := i.(type) {
	case string:
		fmt.Println("The type is string, value = ", v)
	case int:
		fmt.Println("The type is int, value = ", v)
	default:
		fmt.Println("Could not determine the type, value = ", v)
	}

}
