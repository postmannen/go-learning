package main

import (
	"fmt"
	"log"

	"github.com/creack/pty"
)

func main() {
	pt, tt, err := pty.Open()
	if err != nil {
		log.Printf("error: failed to pty.Open\n", err)
	}
	defer pt.Close()
	defer tt.Close()

	fmt.Printf("pty: %v\n", pt.Name())
	fmt.Printf("tty: %v\n", tt.Name())
}
