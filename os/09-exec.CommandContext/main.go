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
		cmd = exec.CommandContext(ctx, "bash", "-c", "tcpdump -nni any")

		// Using cmd.StdoutPipe here so we are continuosly
		// able to read the out put of the command.
		outReader, err := cmd.StdoutPipe()
		if err != nil {
			er := fmt.Errorf("error: methodREQCliCommandCont: cmd.StdoutPipe failed : %v", err)
			log.Printf("error: %v\n", er)
		}
		defer func() {
			err := outReader.Close()
			fmt.Printf(" * closing outReader\n")
			if err != nil {
				log.Printf("error: failed to close errorReader: %v\n", err)
			}
		}()

		errorReader, err := cmd.StderrPipe()
		if err != nil {
			er := fmt.Errorf("error: methodREQCliCommandCont: cmd.StderrPipe failed : %v", err)
			log.Printf("%v\n", er)

			log.Printf("error: %v\n", err)
		}
		defer func() {
			fmt.Printf(" * closing errorReader\n")
			err := errorReader.Close()
			if err != nil {
				log.Printf("error: failed to close outReader: %v\n", err)
			}
		}()

		if err := cmd.Start(); err != nil {
			er := fmt.Errorf("error: methodREQCliCommandCont: cmd.Start failed : %v", err)
			log.Printf("%v\n", er)
		}

		go func() {
			scanner := bufio.NewScanner(errorReader)
			for scanner.Scan() {
				errCh <- scanner.Text()
			}
			fmt.Printf("* DEBUG 4\n")
		}()

		go func() {
			scanner := bufio.NewScanner(outReader)
			for scanner.Scan() {

				text := scanner.Text()
				fmt.Println(text)
				outCh <- []byte(text + "\n")
			}
			fmt.Printf("* DEBUG 3\n")

		}()

		fmt.Printf("* DEBUG 1\n")
		defer func() {
			fmt.Printf("* DEBUG 2\n")
		}()
		<-ctx.Done()

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
