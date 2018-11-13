//The idea here is to test out sending functions as values
// over a channel. Since a functions is also a first class
// citizen like string and int we should be able to that.
package main

import "fmt"

//aFunction is declaration of a type function, with a function
// that takes no input and returns no return values.
type aFunction func()

//Create some functions to put on the channel.
// Since the anonymous func below neither takes any input, or
// return anything, it fullfills the the declaration of the type
// aFunction.
func createFunctions(ch chan aFunction) {
	for i := 0; i < 10; i++ {
		ch <- func() {
			fmt.Println("Value of i = ", i)
		}

	}
	close(ch)
}

func main() {

	ch := make(chan aFunction)
	go createFunctions(ch)

	// Range the channel, and execute all the functions comming.
	for f := range ch {
		f()
	}

}
