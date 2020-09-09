package main

import (
	"fmt"
	"log"

	arg "github.com/postmannen/parrotbebop/conceptTesting/decodingArguments2_makingItAPackageBenchmarking"
)

func main() {
	//The data below should decode
	//bytes 1->4 = a float32,
	//byte 5 = an int8
	//byte 6+ = a string = "Hello There !!!!", it is terminated with a zero at the end.
	//						72 101 108 108 111 32 116 104 101 114 101 32 33 33 33 33 0

	//WORKING FROM HERE
	tmpData := []byte{154, 221, 45, 61, 83, 72, 101, 108, 108, 111, 32, 116, 104, 101, 114, 101, 32, 33, 33, 33, 33, 0, 72, 101, 108, 108, 111, 0, 83}

	// This struct is to simulate the arguments belonging to a command.
	// and we will fill a variable of this type with the arguments read
	// from the byte slice above.
	commandArguments := &struct {
		SomeFloatValue   float32
		SomeIntValue     int8
		SomeStringValue  string
		SomeStringValue2 string
		SomeIntValue2    int8
	}{}

	argDecoder := arg.NewDecoder()

	// Create a map where the key is all the possible commands for the drone,
	// and the values are a function who will decode all the arguments for
	// that specific command
	commandArgumentsMap := make(map[string]func() ([]interface{}, error))
	commandArgumentsMap["cmd1"] = func() ([]interface{}, error) {
		return argDecoder.DecodeArgs(commandArguments, tmpData, &arg.Float, &arg.I8, &arg.Stringx, &arg.Stringx, &arg.I8)
	}

	// Look up an argument function to execute for the command "cmd1"
	fn := commandArgumentsMap["cmd1"]
	argSlice, err := fn()
	if err != nil {
		fmt.Println("error: argumentsToDecode: failed looping over v ", err)
	}

	// Fill the argument struct with the differet values received.
	err = arg.InsertArgValueIntoStruct(commandArguments, argSlice)
	if err != nil {
		log.Printf("error: insertArgValueIntoStruct: %v\n", err)
	}

	fmt.Printf("%+v\n", commandArguments)
}
