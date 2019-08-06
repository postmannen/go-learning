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

// Declare the sizes used for the values received from the drone.

type projectDef uint8 //the first byte
type classDef uint8   // the second byte
type cmdDef uint16    // the 3 and 4th byte, given in little endian.

// command is the representation of a command recevied from the drone.
// The command is received from the drone as a part of an UDP stram
// with the following order.
type command struct {
	project projectDef
	class   classDef
	cmd     cmdDef
}

// project definitions
const ardrone3 projectDef = 1

// class definitions
const pilotingState classDef = 4

// cmd definitions
const attitudeChange cmdDef = 6
const speedChanged cmdDef = 5

// -----------------------The the complete command definitions as a struct----------

// --- add one command for testing

// ardrone3PilotingStateAttitudeChange is a command, and
// it is created so we can have individual decode functions
// for each command definition.
type ardrone3PilotingStateAttitudeChange command

// decode will decode the payload that follows after the command
// part of the udp stream.
func (a ardrone3PilotingStateAttitudeChange) decode() {
	// TODO: Put in the decoding of the variable data part following the command here.
	fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
	fmt.Printf("%+v\n", a)
}

// PilotingStateAttitudeChange are the specification of a single
// command, and the values received from the drone.
var PilotingStateAttitudeChange = ardrone3PilotingStateAttitudeChange{
	project: ardrone3,
	class:   pilotingState,
	cmd:     attitudeChange,
}

// --- Adding one more command for testing

type ardronePilotingStateSpeedChanged command

func (a ardronePilotingStateSpeedChanged) decode() {
	// TODO: Put in the decoding of the variable data part following the command here.
	fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
	fmt.Printf("%+v\n", a)
}

// PilotingStateSpeedChanged defines a speed changed command
var PilotingStateSpeedChanged = ardrone3PilotingStateAttitudeChange{
	project: ardrone3,
	class:   pilotingState,
	cmd:     speedChanged,
}

// decoder is an interface for all types that have a decode method,
// this is being used on the commandMap to map all the values of thetypes
// that start with ardrone3 to it's belonging methods.
// The idea is that we get an incommming udp value from the drone, for for
// example {1,4,6} which is a command, we can then check for that value in
// the index which contains all the commands as structs, and return it's
// decode method to decode the payload.
type decoder interface {
	decode()
}

// commandMap is a map with an index holding all the possible commands for
// the drone, and the value is the method that belongs to that specific
// command type.
var commandMap map[command]decoder

func main() {
	// Make a map to map all the command defenitions to it's method for
	// quick lookup.
	commandMap = make(map[command]decoder)
	commandMap[command(PilotingStateAttitudeChange)] = PilotingStateAttitudeChange
	commandMap[command(PilotingStateSpeedChanged)] = PilotingStateSpeedChanged

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
