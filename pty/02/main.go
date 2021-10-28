package main

import (
	"fmt"
	"log"
	"time"

	"github.com/creack/pty"
)

// Typically you can check the output of what is put into the pty by
// doing a `cat /dev/ttys<x>`

func main() {
	pt, tt, err := pty.Open()
	if err != nil {
		log.Printf("error: failed to pty.Open: %v\n", err)
	}
	defer pt.Close()
	defer tt.Close()

	fmt.Printf("pty: %v\n", pt.Name())
	fmt.Printf("tty: %v\n", tt.Name())

	for i := 0; i < 30; i++ {
		fmt.Fprintf(pt, "test\n")
		time.Sleep(time.Second)
	}
}
