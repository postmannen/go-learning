// Small example for a state machine.
// We create a struct which holds all the states in fields,
// We then create methods on that struct to manipulate the
// state, and in the end the method will return the next
// method to be executed.
// When the last method is received the loop method will
// receive a function value == <nil> and exit back to main,
// and end the program.

package main

import (
	"fmt"
	"strings"
)

type animalParser struct {
	dataSource    []string // The input data
	indexPosition int      // indexPosition for where the position we're working in the dataSource slice.
	names         []string // The processed data
	currentName   string   // The current name from the datasource we're working with
}

// newAnimal will take a []string with animal names as input,
// and return a pointer to an animalParser struct.
func newAnimal(d []string) *animalParser {
	return &animalParser{
		dataSource:    d,
		indexPosition: 0,
	}

}

func (a *animalParser) start() {
	// We need to kickstart the process by reading the first value
	// from the input slice.
	fn := a.readNext()
	for {
		// execute the current function, and put the return value
		// into fn, so we can exexute that on the next iteration.
		fn = fn()

		// If no function was returned, and we received the value
		// <nil> we know that we are done, and can return to main.
		if fn == nil {
			return
		}
	}
}

// animalFunc is a function type which describes a function that
// takes no arguments as input, and returns an animalFunc
type animalFunc func() animalFunc

func (a *animalParser) readNext() animalFunc {
	// Check if we have reached the end of the input slice,
	// and return nil, indicating no more functions to execute.
	if a.indexPosition == len(a.dataSource) {
		return nil
	}

	a.currentName = a.dataSource[a.indexPosition]
	// prepare the index position so the next character will be read
	// next time.
	a.indexPosition++

	return a.title
}

// title will make the first letter of the animal name capitalized.
func (a *animalParser) title() animalFunc {
	a.currentName = strings.Title(a.currentName)

	return a.currentAnimalDone
}

// currentAnimalDone will append the current animal to the output slice,
// and return a function to read the next value from the input
func (a *animalParser) currentAnimalDone() animalFunc {
	a.names = append(a.names, a.currentName)

	return a.readNext
}

func main() {
	// Create som input animal data
	animals := []string{"rhino", "sparrow", "snake", "ant", "fly", "hippo"}

	// Create a new variable of type animal, and prepare it with the
	// input data.
	a := newAnimal(animals)

	// Start the state machine, and work on the input data.
	a.start()

	fmt.Printf("Input animals = %+v\n", animals)
	fmt.Printf("Output animals = %+v\n", a.names)

}
