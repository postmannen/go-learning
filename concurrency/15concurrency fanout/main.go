package main

import (
	"fmt"
	"time"
)

type car struct {
	name     string
	duration time.Duration
}

func main() {
	cars := 20
	carCh := make(chan car, cars)

	//Create as many Go routines as there are cars,
	//and give each car a name
	for i := 0; i < 20; i++ {
		go func(nr int) {
			aCar := car{
				name: fmt.Sprintf("carname%v", nr),
			}
			carCh <- aCar
		}(i)
	}

	//Read all the items on the channel.
	for cars > 0 {
		c := <-carCh
		fmt.Println("Length of channel = ", len(carCh))
		fmt.Printf("The name of the car = %v\n", c.name)
		cars--
	}

}
