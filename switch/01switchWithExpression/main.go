package main

import (
	"fmt"
)

type myStruct struct {
	counter int
	name    string
}

//test returns prepared variable of type myStruct
func test() myStruct {
	return myStruct{
		counter: 10,
		name:    "apekatt",
	}
}

func main() {
	//First we run the test function and assigns it to the variable a.
	//Then we can to the actual switch on one of the fields of a (a.counter)
	switch a := test(); a.counter {
	case 10:
		fmt.Println("The counter was ", a.counter)
	default:
		fmt.Println("The counter was something else")
	}
}
