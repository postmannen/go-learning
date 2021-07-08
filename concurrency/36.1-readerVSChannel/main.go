// The point of the test is to have something that simulates a stream
// of data, here simulated as the data generator.
// We also wants to start 5 go routines who reads that stream of data.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	"log"
)

type dataType struct {
	nr int
}

type dataGenerator struct{}

// Will generate a number and put into nr.
func (d dataGenerator) generate(ch chan dataType) error {
	for {
		ch <- dataType{
			nr: rand.Intn(100),
		}

		// Simulate that we could return an error.
		var err error
		if err != nil {
			return err
		}
	}
}

func main() {
	ch := make(chan dataType)

	d := dataGenerator{}
	go func() {
		err := d.generate(ch)
		if err != nil {
			log.Printf("error: dataGenerator: %v\n", err)
			os.Exit(1)
		}
	}()

	routines := 5
	var wg sync.WaitGroup
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			v := <-ch
			fmt.Printf("%v\n", v.nr)
			wg.Done()
		}()
	}

	wg.Wait()
}
