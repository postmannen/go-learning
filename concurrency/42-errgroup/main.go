package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	{
		g := &errgroup.Group{}

		g.Go(func() error {
			return fmt.Errorf("an error happened")
		})

		g.Go(func() error {
			time.Sleep(time.Second * 5)
			fmt.Printf("message from another....\n")
			return fmt.Errorf("another error happened")
		})

		if err := g.Wait(); err != nil {
			log.Printf("error: %v\n", err)
		}
	}

	{
		g := &errgroup.Group{}

		g.Go(func() error {
			return fmt.Errorf("yet another error happened")
		})

		if err := g.Wait(); err != nil {
			log.Printf("error: %v\n", err)
		}

	}
}
