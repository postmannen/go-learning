/*
Function type, and function who returns a function.
 NB: The important thing to remember here is that it is the returned function
 functions signature that have to match the myFunc type, which is being used
 as input in functionExecutor. The called function can look whatever way it wants,
 just as long as the returned function matches to myFunc signature.
*/

package main

import "fmt"

//myFunc is a function type with no input and no return values.
type myFunc func()

type person struct {
	name string
}

func (p *person) setName(s string) myFunc {
	return func() {
		p.name = s
	}
}

//functionExecutor takes a myFunc as it input argument.
func functionExecutor(m myFunc) {
	m()
}

func main() {
	p1 := &person{}
	functionExecutor(p1.setName("Harald"))

	fmt.Println("p1 now contains : ", p1)

}
