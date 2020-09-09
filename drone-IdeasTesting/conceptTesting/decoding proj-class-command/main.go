/*
The UDP stream received from the drone specifying what command
to use will look like this:
[]byte{1, 4, 6, 0, 66, 170, 134, 60, 21, 152, 181, 189, 131, 170, 224, 191}
where the first 3 values {1,4,06} is a reference to what this command is,
and the following bytes are the payload (arguments) for that command.

What needs to be solved ?
Since all the various commands received from the drone will be made as it's
own type in the code with a decode method attached, we need to be able to look
up what decode method to use for a specific command definition.

This test is a concept to create a map to decode all the commands received,
and find the corresponding decode method to execute and decode the arguments
in the payload.
*/
package main

import (
	"fmt"
)

func main() {

	// Create a command received from the drone.
	{
		c := command{
			project: 1,
			class:   4,
			cmd:     6,
		}

		// Check if the command value exists in the map, and if it does run
		// it's decode method which is stored as a value in the map.
		v, ok := commandMap[c]
		if ok {
			fmt.Printf("The type of v inside the if check = %T\n---\n", v)
			v.decode()
		}
	}
	// -------------------------------- test 2
	fmt.Println("-----------------------------")
	// Create a command received from the drone.
	{
		c := command{
			project: 1,
			class:   4,
			cmd:     5,
		}

		// Check if the command value exists in the map, and if it does run
		// it's decode method which is stored as a value in the map.
		v, ok := commandMap[c]
		if ok {
			fmt.Printf("The type of v inside the if check = %T\n---\n", v)
			v.decode()
		}
	}
}
