package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

func main() {
	f := func(context.Context) error {
		fmt.Printf("test\n")
		return nil
	}

	var wg sync.WaitGroup
	ctx, _ := context.WithCancel(context.Background())

	wg.Add(1)

	go func() {
		err := f(ctx)
		if err != nil {
			log.Printf("error: problem with function to execute: %v\n", err)
		}

		wg.Done()
	}()

	wg.Wait()
}
