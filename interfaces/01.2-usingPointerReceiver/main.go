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
func (h *horse) run() {
	fmt.Println("Running!")
}

// acce
func doWithRunner(r runner) {
	r.run()
}

func main() {
	// since the there is a pointer receiver on the method satisfying the
	// interface, we have to use pointer to values when we want to assign
	// values to satisfy an argument of the interface type.

	// A value type is not allowed when there is a pointer receiver on the
	// method satisfying the interface type,
	// so the value below is not allowed
	//
	// var i1 runner
	// i1 = horse{name: "blitz"}
	// fmt.Printf("%#v", i1)

	var i2 runner
	i2 = &horse{name: "light"}
	doWithRunner(i2)
}
