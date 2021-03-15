// Test with
// go run main.go -startCLISubscriber="a,b,c"
//
// Example using a string flag type, and parsing the
// comma separated values in the string, and then
// filling those values into a struct.
//
// The startsub struct need a "Set" and a "String"
// method to be of a type to be accepted by the value
// interface type.
//
// type Value interface {
//     String() string
//     Set(string) error
// }

package main

import (
	"flag"
	"fmt"
	"strings"
)

// Our type who we are going to fill with values.
type startsub struct {
	value string
	run   bool
	nodes []string
}

// The needed String method to fulfill the interface.
func (f *startsub) String() string {
	return ""
}

// The needed Set method to fulfill the interface.
// If f.value where for example a slice, we could
// use the Set method to append flag values.
func (f *startsub) Set(s string) error {
	f.value = s
	return nil
}

// Custom function to parse the comma separated input
// string into the struct values.
func (f *startsub) Parse() error {
	if len(f.value) == 0 {
		return nil
	}

	fv := f.value
	sp := strings.Split(fv, ",")
	f.run = true
	f.nodes = sp
	return nil
}

func main() {
	var startCLISubscriber startsub

	flag.Var(&startCLISubscriber, "startCLISubscriber", "enter value")
	flag.Parse()

	startCLISubscriber.Parse()

	if startCLISubscriber.run {
		fmt.Printf("%#v\n", startCLISubscriber.nodes)
	}
}
