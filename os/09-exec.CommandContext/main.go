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

	go func() {
		cmd := exec.CommandContext(ctx, "bash", "-c", "timeout 2 tcpdump -nni any > dump.txt")

		// Using cmd.StdoutPipe here so we are continuosly
		// able to read the out put of the command.
		outReader, err := cmd.StdoutPipe()
		if err != nil {
			er := fmt.Errorf("error: methodREQCliCommandCont: cmd.StdoutPipe failed : %v", err)
			log.Printf("error: %v\n", er)
		}

		ErrorReader, err := cmd.StderrPipe()
		if err != nil {
			er := fmt.Errorf("error: methodREQCliCommandCont: cmd.StderrPipe failed : %v", err)
			log.Printf("%v\n", er)

			log.Printf("error: %v\n", err)
		}

		if err := cmd.Start(); err != nil {
			er := fmt.Errorf("error: methodREQCliCommandCont: cmd.Start failed : %v", err)

			log.Printf("%v\n", er)

		}

		// Also send error messages that might happen during the time
		// cmd.Start runs.
		// Putting the scanner.Text value on a channel so we can make
		// the scanner non-blocking, and also check context cancelation.
		go func() {
			errCh := make(chan string, 1)

			scanner := bufio.NewScanner(ErrorReader)
			for scanner.Scan() {
				select {
				case errCh <- scanner.Text():
					er := fmt.Errorf("error: methodREQCliCommandCont: cmd.Start failed : %v, error_output: %v", err, <-errCh)

					log.Printf("%v\n", er)
				case <-ctx.Done():
					return
				}
			}
		}()

		scanner := bufio.NewScanner(outReader)
		for scanner.Scan() {
			select {
			case outCh <- []byte(scanner.Text() + "\n"):
			case <-ctx.Done():
				return
			}

		}

	}()

	// Check if context timer or command output were received.
	for {
		select {
		case <-ctx.Done():
			cancel()
			er := fmt.Errorf("info: methodREQCliCommandCont: method timeout reached, canceling: methodArgs")
			log.Printf("%v\n", er)
			return
		case out := <-outCh:
			// Prepare and queue for sending a new message with the output
			// of the action executed.
			fmt.Printf("out: %v", string(out))
		}
	}
}
