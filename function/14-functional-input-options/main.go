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

func printAnInt(i int) myFunc {
	return func() {
		fmt.Println("The int = ", i)
	}
}

//functionExecutor takes a myFunc as it input argument.
func functionExecutor(m myFunc) {
	m()
}

func main() {
	functionExecutor(printAnInt(20))

}
