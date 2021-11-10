package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", "uname -a")

	out, _ := cmd.StdoutPipe()
	//cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Printf("error: failed cmd.Run: %v\n", err)
	}
	defer cmd.Wait()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer fmt.Println()
		for {
			b := make([]byte, 1)
			_, err := out.Read(b)
			if err == io.EOF {
				break
			}
			fmt.Printf("%v", string(b))
		}
		wg.Done()
	}()

	wg.Wait()

}
