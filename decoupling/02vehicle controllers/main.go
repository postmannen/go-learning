package main

import (
	"fmt"
	"log"
	"math/rand"
)

//tractorBus simulates a bus on a tractor controller
type tractorBus struct {
	bus chan int
}

func newTractorBus() *tractorBus {
	tb := make(chan int)
	return &tractorBus{
		bus: tb,
	}
}

//Start will make the tractor bus producing values on the channel
func (t *tractorBus) Start() {
	go func() {
		for {
			t.bus <- rand.Intn(10)
		}
	}()

}

//Read will read 1 value from the channel
func (t *tractorBus) Read() (int, error) {
	d := <-t.bus
	return d, nil
}

// ==========================================================

// ==========================================================

func main() {
	tractor := newTractorBus()
	tractor.Start()
	d, err := tractor.Read()
	if err != nil {
		log.Println("Error: Read tractor: ", err)
	}

	fmt.Println("Read returned : ", d)

}
