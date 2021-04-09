package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type OpCommandRaw struct {
	Cmd    string
	OpArgs json.RawMessage
}

type OpStart struct {
	AllowedNodes []string
}

type OpStop struct {
	Node string
}

func main() {
	var err error

	var j = []byte(`[
	{"Cmd": "OpStart", "OpArgs": {"AllowedNodes":["host1","host2"]}},
	{"Cmd": "OpStop", "OpArgs": {"Node": "host3"}}
	]`)

	// First Unmarshal the outer structure that will give use the Cmd, and
	// a field called OpArgs containing the raw json data that will represent
	// either the OpStart or OpStop type structs.
	var opCmds []OpCommandRaw
	err = json.Unmarshal(j, &opCmds)
	if err != nil {
		log.Printf("error: unmarshal: %v\n", err)
	}
	fmt.Printf("type=%T, %v\n\n", opCmds, opCmds)

	// Range over the slice we got after unmarshalling earlier, and unmarshal
	// the json.RawMessage field called OpArgs.
	for _, oc := range opCmds {
		// Dst interface is the generic type to Unmarshal OpArgs into, and we will
		// set the type it should contain depending on the value specified in Cmd.
		var dst interface{}

		switch oc.Cmd {
		case "OpStart":
			// Set the interface type dst to &OpStart.
			dst = &OpStart{}

			err := json.Unmarshal(oc.OpArgs, &dst)
			if err != nil {
				log.Printf("error: unmarshal: %v\n", err)
			}

			// Print it out, and also assert it into the correct non pointer value.
			fmt.Printf("type=%T, content=%#v\n\n", *dst.(*OpStart), *dst.(*OpStart))

		case "OpStop":
			// Set the interface type dst to &OpStart.
			dst = &OpStop{}

			err := json.Unmarshal(oc.OpArgs, &dst)
			if err != nil {
				log.Printf("error: unmarshal: %v\n", err)
			}

			// Print it out, and also assert it into the correct non pointer value.
			fmt.Printf("type=%T, content=%#v\n\n", dst, dst)
		}

	}

}
