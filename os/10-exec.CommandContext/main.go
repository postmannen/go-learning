package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(3))

	outCh := make(chan []byte)
	errCh := make(chan string, 1)
	var cmd *exec.Cmd

	go func() {
		cmd = exec.CommandContext(ctx, "bash", "-c", "tcpdump -nni any -l")

		// Using cmd.StdoutPipe here so we are continuosly
		// able to read the out put of the command.
		outReader, err := cmd.StdoutPipe()
		if err != nil {
			log.Printf("error: cmd.StdoutPipe failed:%v\n", err)
		}

		errorReader, err := cmd.StderrPipe()
		if err != nil {
			log.Printf("error: cmd.StderrPipe failed:%v\n", err)

			log.Printf("error: %v\n", err)
		}

		if err := cmd.Start(); err != nil {
			log.Printf("error: cmd.Start failed:%v\n", err)
		}

		go func() {
			buf := bufio.NewReader(errorReader)
			for {
				s, err := buf.ReadString('\n')
				if err != nil {
					log.Printf("error: reading errorReader: %v\n", err)
					break
				}
				errCh <- s
			}
			fmt.Printf("* DEBUG 4\n")
		}()

		go func() {
			buf := bufio.NewReader(outReader)
			for {
				s, err := buf.ReadString('\n')
				fmt.Println("apekatt: ", s)
				if err != nil {
					log.Printf("error: reading outReader: %v\n", err)
					break
				}
				errCh <- s
			}
			fmt.Printf("* DEBUG 3\n")

		}()

		fmt.Printf("* DEBUG 1\n")
		defer func() {
			fmt.Printf("* DEBUG 2\n")
		}()

		err = cmd.Wait()
		if err != nil {
			log.Printf("error: cmd.Wait: %v\n", err)
		}
	}()

	// Check if context timer or command output were received.
	for {
		select {
		case <-ctx.Done():
			cancel()
			er := fmt.Errorf("info: methodREQCliCommandCont: method timeout reached, canceling: methodArgs")
			log.Printf("%v\n", er)
			goto end
		case out := <-outCh:
			// Prepare and queue for sending a new message with the output
			// of the action executed.
			fmt.Printf("out: %v", string(out))
		case errOut := <-errCh:
			er := fmt.Errorf("error: methodREQCliCommandCont: cmd.Start failed : error_output: %v", errOut)

			log.Printf("%v\n", er)
		}
	}
end:
}
