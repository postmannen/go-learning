package main

import (
	"fmt"
	"log"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		log.Println("failed: termbox init: ", err)
	}

	defer termbox.Close()

	keyCheck := true

	for keyCheck {
		//Check for any new event
		switch event := termbox.PollEvent(); event.Type {
		//If the event is a key press
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyEsc:
				log.Println("exiting")
				//Break out of for loop
				keyCheck = false
			case termbox.KeyArrowUp:
				fmt.Println("Drone up")
			case termbox.KeyArrowDown:
				fmt.Println("Drone down")
			}

		}
	}

}
