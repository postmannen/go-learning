package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type event struct {
	eventType eventType
}

type eventType int

const (
	etPrint eventType = iota
	etExit
	etDone
)

var wg sync.WaitGroup

func main() {

	errCh := make(chan error, 1)
	eventCh := make(chan event, 1)
	ctx := context.Background()

	wg.Add(1)
	go func() {
		errCh <- readEvents(ctx, eventCh)
	}()

	eventCh <- event{eventType: etPrint}

	wg.Wait()

	err := <-errCh
	log.Printf("%v\n", err)

}

func readEvents(ctx context.Context, eventCh chan event) error {
	defer wg.Done()
	for {
		select {
		case e := <-eventCh:
			switch e.eventType {
			case etPrint:
				fmt.Printf("info: got event: %v\n", e)
				eventCh <- event{eventType: etDone}
			case etDone:
				fmt.Printf("info: got event: %v\n", e)
				return fmt.Errorf("got etDone")
			}
		case <-ctx.Done():
			return fmt.Errorf("info: got ctx.Done")
		}
	}
}
