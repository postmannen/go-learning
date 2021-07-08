package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type dataType struct {
	nr int
}

type dataGenerator struct {
	ch chan dataType
}

func newDataGenerator() *dataGenerator {
	dg := dataGenerator{
		ch: make(chan dataType),
	}
	return &dg
}

func (d *dataGenerator) generate() {
	go func() {
		for {
			dt := dataType{
				nr: rand.Intn(100),
			}
			d.ch <- dt
		}
	}()
}

// Will generate a number and put into nr.
func (d *dataGenerator) dRead(dt *dataType) error {
	*dt = <-d.ch
	return nil
}

func main() {
	dg := newDataGenerator()
	dg.generate()

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
