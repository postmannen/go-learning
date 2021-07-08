package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type dataType struct {
	nr int
}

// dataGenerator enulates f.ex. something that reads a socket,
// prepares those data read, and can be read elsewhere via the
// dRead method.
type dataGenerator struct {
	ch chan dataType
}

func newDataGenerator() *dataGenerator {
	dg := dataGenerator{
		ch: make(chan dataType),
	}
	return &dg
}

// run will start the process of reading f.ex. a fictive socket.
func (d *dataGenerator) run(f func() dataType) error {
	for {
		dt := f()
		d.ch <- dt

		// simulate that we can get catch an error when generating
		// data, and return if there is a problem.
		var err error
		if err != nil {
			return err
		}
	}
}

// Will read the next available dataType into dt.
func (d *dataGenerator) dRead(dt *dataType) error {
	*dt = <-d.ch
	return nil
}

func main() {

	generatefunc := func() dataType {
		return dataType{
			nr: rand.Intn(100),
		}
	}

	dg := newDataGenerator()
	go func() {
		err := dg.run(generatefunc)
		if err != nil {
			return
		}
	}()

	var dt dataType

	routines := 5
	var wg sync.WaitGroup
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			dg.dRead(&dt)
			fmt.Printf("%v\n", dt.nr)
			wg.Done()
		}()
	}
	wg.Wait()

}
