package main

import (
	"fmt"
	"log"
	"math/rand"
)

// ==========================================================

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
//carBus simulates a bus on a tractor controller
type carBus struct {
	bus chan int
}

func newCarBus() *carBus {
	cb := make(chan int)
	return &carBus{
		bus: cb,
	}
}

//Start will make the tractor bus producing values on the channel
func (c *carBus) Start() {
	go func() {
		for {
			c.bus <- rand.Intn(10)
		}
	}()

}

//Read will read 1 value from the channel
func (c *carBus) Read() (int, error) {
	d := <-c.bus
	return d, nil
}

// ==========================================================

//Reader is a reader interface type for all vehicles
type Reader interface {
	Read() (int, error)
}

// ==========================================================

func readToScreen(v Reader, amount int) error {
	tractorData, err := v.Read()
	if err != nil {
		log.Println("Error: Read tractor: ", err)
	}

	fmt.Println("Read tractor returned : ", tractorData)
	return nil
}

func main() {
	tractor := newTractorBus()
	tractor.Start()
	readToScreen(tractor, 5)

	car := newCarBus()
	car.Start()
	readToScreen(car, 5)

}
