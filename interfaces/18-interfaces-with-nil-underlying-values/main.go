package main

import "fmt"

//I is an interface with the description av a method M()
type I interface {
	M()
}

//T is a struct type
type T struct {
	S string
}

//M is a method for a type
func (t *T) M() {
	fmt.Println("---------------------Calling method for *T")
	if t == nil {
		fmt.Println("The content of the receiver t = <nil>")
		return
	}
	fmt.Println("Now the receiver t holds something, and t.S = ", t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Println("---------------------------Calling 'describe', taking interface I as input----------------------------")
	fmt.Printf("describe takes the interface I, and it now contains = (%v, %T)\n", i, i)
}
