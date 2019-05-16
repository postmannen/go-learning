/*
	Using sync.Once to make sure a function is only run once.
	Spinning up several go routines of the actual checkout to simulate a hickup
	that causes the checkout of the same order to be tried out at the same time.
*/
package main

import (
	"fmt"
	"sync"
)

type order struct {
	id         int
	once       sync.Once
	checkedOut bool
}

//Checkout will do the actual checkout of the order
// Using sync.Once to ensure the actual checkout is only done once.
func (c *order) Checkout() {
	c.once.Do(
		func() {
			fmt.Printf("*********************** checking out ID = %v **********************\n", c.id)
			c.checkedOut = true
		},
	)
}

//PrepareCheckout will do final preparations and checks before the actual checkout.
func (c *order) PrepareCheckout() error {
	//This one is actually a little interesting.
	// Since c.checkedOut is set to true after a checkout is done, it can occur that another
	// go routine will reach this point and try to check if c.checkedOut is false, and start
	// the actual checkout() function again. But since we have a sync.Once in the Checkout() function
	// it is not checked out a second time (as the output shows).
	if c.checkedOut {
		return fmt.Errorf("Order with ID=%v is allready checked out.\n", c.id)
	}
	c.Checkout()
	return nil
}

func main() {
	//Here we simulate the checkout of the order
	// We start the checkout in a go routine to simulate a system gone bananas.
	orderOne := &order{id: 1}
	const checkoutsToSimulate = 10

	var wg sync.WaitGroup
	wg.Add(checkoutsToSimulate)
	for i := 1; i <= checkoutsToSimulate; i++ {
		go func(i int) {
			fmt.Println("-----------------------------------------------------------")
			fmt.Printf("Running the Prepare checkout for the %v time\n", i)
			err := orderOne.PrepareCheckout()
			if err != nil {
				fmt.Printf("checkout run = %v, err = %v\n", i, err)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
