package main

import (
	"context"
	"testing"
)

func TestPut(t *testing.T) {
	cM := newCMap()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//var wg sync.WaitGroup

	//wg.Add(1)
	go func() {
		cM.run(ctx)
		//wg.Done()
	}()

	cM.put(keyValue{k: 1, v: "a"})

	//wg.Wait()
}
