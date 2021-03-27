// When using exec.CommandContext and executing a command with
// "bash -c" it seems that the timeout of of the context do
// not cancel the
package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	outCh := make(chan []byte)

	go func() {
		cmd := exec.CommandContext(ctx, "bash", "-c", "sleep 5 && ls -l")
		out, err := cmd.Output()
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		outCh <- out
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("done received\n")
	case o := <-outCh:
		log.Printf("o: %s\n", o)
	}

}
