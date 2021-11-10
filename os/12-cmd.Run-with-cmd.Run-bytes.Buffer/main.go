package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "-c", "uname -a")

	var out bytes.Buffer
	// var stderr bytes.Buffer
	cmd.Stdout = &out
	//cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("error: failed cmd.Run: %v\n", err)
	}

	fmt.Printf(" * cmd: %v\n", cmd)

	defer fmt.Println()
	for _, b := range out.Bytes() {
		fmt.Printf("%v", string(b))
	}

}
