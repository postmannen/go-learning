package main

import (
	"fmt"
)

type runner interface {
	run()
}

type horse struct {
	name string
}

// run is defined with a value receiver.
func (h horse) run() {
	fmt.Println("Running!")
}

// acce
func doWithRunner(r runner) {
	r.run()
}

func main() {
	// since the there is a value receiver, and not a pointer receiver on
	// the run method, both values of pointer and value types are allowed
	// and will be accepted by the interface argument.

	var i1 runner
	i1 = horse{name: "blitz"}
	doWithRunner(i1)

	var i2 runner
	i2 = &horse{name: "light"}
	doWithRunner(i2)
}
