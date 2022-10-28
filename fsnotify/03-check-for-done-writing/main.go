package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gopkg.in/fsnotify.v1"
)

type configuration struct {
	waitWrite int
}

func newConfiguration() *configuration {
	c := configuration{}

	flag.IntVar(&c.waitWrite, "waitWrite", 2000, "time in milliseconds for how long to wait for write operations before we assume that the writing is done")
	flag.Parse()

	return &c
}

// Start listening for events.
func eventAllListener(ctx context.Context, watcher *fsnotify.Watcher, mu *sync.Mutex, fMap map[string]chan fsnotify.Event, eventCh chan chan fsnotify.Event) error {

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return fmt.Errorf("error: got !ok while receiving <-watcher.Events")
			}

			// If a Create event is received, create the write channel that
			// will be used to communicate the Write events for a specific file.
			// We send the created channel to the listening go routine, where
			// it will start up a new go routine specific for the created file,
			// and listen for Write events to that file.
			if event.Op == fsnotify.Create {
				log.Printf("info: got create event: %v\n", event)
				ch := make(chan fsnotify.Event)

				mu.Lock()
				fMap[event.Name] = ch
				mu.Unlock()

				eventCh <- ch

				// Also send a message to start up the go routine that
				// will receive the Write events.
			}

			// For Write events we look up the correct channel to use in the map,
			// and send the events on the channel to the go routine that handles
			// the Write event for that file.
			if event.Op == fsnotify.Write {

				mu.Lock()
				ch, ok := fMap[event.Name]
				mu.Unlock()
				if !ok {
					log.Printf("info: received Write, but no entry for %v in map, so we probably never got a create event for this file\n", event.Name)
					continue
				}
				ch <- event

			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return fmt.Errorf("error: watcher error: %v", err)
			}
			log.Println("error:", err)

		case <-ctx.Done():
			return nil
		}
	}

}

func handleWriteEvents(ctx context.Context, eventCh chan chan fsnotify.Event, configuration configuration, mu *sync.Mutex, fMap map[string]chan fsnotify.Event) error {

	for {
		select {
		// Get the channel that will be used to receive the Write events.
		case ch := <-eventCh:

			// Since we've got here, it means that a Create event have for
			// file have happened earlier. Start a go routine to handle
			// the Write events for that specific file.
			go func(ch chan fsnotify.Event) {
				ticker := time.NewTicker(time.Millisecond * time.Duration(configuration.waitWrite))

				// Create a variable to hold the last event to use when the ticker is received.
				var ev fsnotify.Event

				for {

					select {
					case event := <-ch:
						ev = event
						// log.Printf("info: got event, reset'ing wait timer for more writes: %v\n", event)
						// Reset the ticker timer
						ticker.Reset(time.Millisecond * time.Duration(configuration.waitWrite))

					case <-ticker.C:
						log.Printf("info: reached max write wait time, deleting entry in map for : %v\n", ev.Name)

						mu.Lock()
						delete(fMap, ev.Name)
						mu.Unlock()
						log.Printf("info: Writing file done, and successfully deleted entry in map for : %v\n", ev.Name)

						close(ch)

						return

					case <-ctx.Done():
						return
					}
				}
			}(ch)

		case <-ctx.Done():
			return nil
		}

	}

}

func main() {
	conf := newConfiguration()

	// We need a map that describes the files currently being worked on,
	// and the channel to the go routine that receives the Write events
	// for each file.
	fMap := make(map[string]chan fsnotify.Event)
	var mu sync.Mutex

	// Handle ctrl+c.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	eventCh := make(chan chan fsnotify.Event)
	errorCh := make(chan error)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("error: creating watcher: %v\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := eventAllListener(ctx, watcher, &mu, fMap, eventCh)
		if err != nil {
			errorCh <- err
		}
	}()

	// Add the path to watch.
	err = watcher.Add("./")
	if err != nil {
		log.Fatal(err)
	}

	// Handle all the Write events here.
	go func() {
		err := handleWriteEvents(ctx, eventCh, *conf, &mu, fMap)
		if err != nil {
			errorCh <- err
		}
	}()

	select {
	case <-sigCh:
		log.Printf("info: got ctrl+c, exiting\n")
	case <-errorCh:
		log.Printf("%v\n", err)
		os.Exit(1)
	}

}
