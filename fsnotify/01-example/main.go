package main

import (
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

func waitUntilFind(filename string) error {
	for {
		time.Sleep(1 * time.Second)
		_, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return err
			}
		}
		break
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("error")
	}
	filename := os.Args[1]
	err := waitUntilFind(filename)
	if err != nil {
		log.Fatalln(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	err = watcher.Add(filename)
	if err != nil {
		log.Fatalln(err)
	}

	renameCh := make(chan bool)
	removeCh := make(chan bool)
	errCh := make(chan error)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				switch {
				case event.Op&fsnotify.Write == fsnotify.Write:
					log.Printf("Write:  %s: %s", event.Op, event.Name)
				case event.Op&fsnotify.Create == fsnotify.Create:
					log.Printf("Create: %s: %s", event.Op, event.Name)
				case event.Op&fsnotify.Remove == fsnotify.Remove:
					log.Printf("Remove: %s: %s", event.Op, event.Name)
					removeCh <- true
				case event.Op&fsnotify.Rename == fsnotify.Rename:
					log.Printf("Rename: %s: %s", event.Op, event.Name)
					renameCh <- true
				case event.Op&fsnotify.Chmod == fsnotify.Chmod:
					log.Printf("Chmod:  %s: %s", event.Op, event.Name)
				}
			case err := <-watcher.Errors:
				errCh <- err
			}
		}
	}()

	go func() {
		for {
			select {
			case <-renameCh:
				err = waitUntilFind(filename)
				if err != nil {
					log.Fatalln(err)
				}
				err = watcher.Add(filename)
				if err != nil {
					log.Fatalln(err)
				}
			case <-removeCh:
				err = waitUntilFind(filename)
				if err != nil {
					log.Fatalln(err)
				}
				err = watcher.Add(filename)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}()

	log.Fatalln(<-errCh)
}
