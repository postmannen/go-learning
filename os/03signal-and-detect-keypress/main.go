package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

func main() {

	//Using bufio.NewScanner
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			key := fmt.Sprintln(scanner.Text())
			fmt.Printf("--NewScanner-- %v\n", key)
			fmt.Printf("--NewScanner-- %v\n", []byte(key))
		}
	}()

	//Using bufio.NewReader
	go func() {
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("--NewReader--- %d \n", char)
	}()

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	//Block until we receive a signal
	s := <-sigCh
	fmt.Printf("Got signal %v, which is of type %T\n", s, s)
}
