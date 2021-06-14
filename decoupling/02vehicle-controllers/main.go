package main

import (
	"fmt"
	"log"
	"math/rand"
)

// ==========================================================

//tractorBus simulates a bus on a tractor controller
type tractorBus struct {
	name string
	bus  chan int
}

func newTractorBus(name string) *tractorBus {
	tb := make(chan int)
	return &tractorBus{
		name: name,
		bus:  tb,
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
	name string
	bus  chan int
}

func newCarBus(name string) *carBus {
	cb := make(chan int)
	return &carBus{
		name: name,
		bus:  cb,
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
	for i := 0; i < amount; i++ {
		tractorData, err := v.Read()
		if err != nil {
			log.Printf("Error: Read %T: %v \n ", v, err)
		}

		fmt.Printf("Read %T returned : %v\n", v, tractorData)
	}
	return nil
}

// ==========================================================

func main() {
	tractor := newTractorBus("John Deere")
	tractor.Start()
	readToScreen(tractor, 5)

	car := newCarBus("Nissan")
	car.Start()
	readToScreen(car, 5)

}
