package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type operation struct {
	OpCmd string          `json:"opCmd"`
	OpArg json.RawMessage `json:"opArg"`
}

type opStartProc struct {
	AllowedNodes []string `json:"allowedNodes"`
}

type opStopProc struct {
	Node string `json:"node"`
}

func main() {
	var err error

	var j = []byte(`{"OpCmd": "startProc", "opArg": {"allowedNodes":["host1","host2"]}}`)
	//var j = []byte(`{"opCmd": "stopProc", "opArg": {"Node": "host3"}}`)

	// First Unmarshal the outer structure that will give use the Cmd, and
	// a field called OpArgs containing the raw json data that will represent
	// either the OpStart or OpStop type structs.
	var opCmd operation
	err = json.Unmarshal(j, &opCmd)
	if err != nil {
		log.Printf("error: unmarshal: %v\n", err)
	}
	fmt.Printf("type=%T, %v\n\n", opCmd, opCmd)

	// unmarshal the json.RawMessage field called OpArgs.
	//
	// Dst interface is the generic type to Unmarshal OpArgs into, and we will
	// set the type it should contain depending on the value specified in Cmd.
	var dst interface{}

	switch opCmd.OpCmd {
	case "startProc":
		// Set the interface type dst to &OpStart.
		dst = &opStartProc{}

		err := json.Unmarshal(opCmd.OpArg, &dst)
		if err != nil {
			log.Printf("error: outer unmarshal: %v\n", err)
		}

		// Print it out, and also assert it into the correct non pointer value.
		fmt.Printf("type=%T, content=%#v\n\n", *dst.(*opStartProc), *dst.(*opStartProc))

	case "stopProc":
		// Set the interface type dst to &OpStart.
		dst = &opStopProc{}

		err := json.Unmarshal(opCmd.OpArg, &dst)
		if err != nil {
			log.Printf("error: inner unmarshal: %v\n", err)
		}

		// Print it out, and also assert it into the correct non pointer value.
		fmt.Printf("type=%T, content=%#v\n\n", dst, dst)
	}

}
