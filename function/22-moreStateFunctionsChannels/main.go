// Small example for a state machine.
// We create a struct which holds all the states in fields,
// We then create methods on that struct to manipulate the
// state, and in the end each method will report the next
// method to be executed on the channel.
// When the last method is done, the channel will be closed,
// the for loop ranging over the channel will stop,  and we
// exit back to main and end the program.

package main

import (
	"fmt"
	"strings"
)

type animalParser struct {
	dataSource    []string    // The input data
	indexPosition int         // indexPosition for where the position we're working in the dataSource slice.
	names         []string    // The processed data
	currentName   string      // The current name from the datasource we're working with
	funcCh        chan func() // channel used to send the next function to be executed.
}

// newAnimal will take a []string with animal names as input,
// and return a pointer to an animalParser struct.
func newAnimal(d []string) *animalParser {
	return &animalParser{
		dataSource:    d,
		indexPosition: 0,
		// Set buffer to one so we can put one value on the channel
		// before have anything reading from it to avoid deadlock.
		funcCh: make(chan func(), 1),
	}

}

func (a *animalParser) start() {
	// We need to kickstart the process by reading the first value
	// from the input slice. This is also why we need a channel
	// with the buffer of 1 since nothing is yet reading from that
	// channel. If the buffer was set to zero it would deadlock here.
	a.funcCh <- a.readNext

	for fn := range a.funcCh {
		// execute the current function, and put the return value
		// into fn, so we can exexute that on the next iteration.
		go fn()

		//// If no function was returned, and we received the value
		//// <nil> we know that we are done, and can return to main.
		//if fn == nil {
		//	fmt.Println("*** fn = <nil>")
		//	break
		//} else {
		//	fmt.Println("*** fn != <nil>")
		//}
	}
}

func (a *animalParser) readNext() {
	// Check if we have reached the end of the input slice,
	// and close the channel, indicating no more functions to execute.
	if a.indexPosition == len(a.dataSource) {
		close(a.funcCh)
		return
	}

	a.currentName = a.dataSource[a.indexPosition]

	// prepare the index position so the next character will be read
	// next time.
	a.indexPosition++

	a.funcCh <- a.title
}

// title will make the first letter of the animal name capitalized.
func (a *animalParser) title() {
	a.currentName = strings.Title(a.currentName)

	a.funcCh <- a.currentAnimalDone
}

// currentAnimalDone will append the current animal to the output slice,
// and return a function to read the next value from the input
func (a *animalParser) currentAnimalDone() {
	a.names = append(a.names, a.currentName)

	a.funcCh <- a.readNext
}

func main() {
	// Create some input animal data
	animals := []string{"rhino", "sparrow", "snake", "ant", "fly", "hippo"}

	// Create a new variable of type animal, and prepare it with the
	// input data.
	a := newAnimal(animals)

	// Start the state machine, and work on the input data.
	a.start()

	fmt.Printf("Input animals = %+v\n", animals)
	fmt.Printf("Output animals = %+v\n", a.names)

}
