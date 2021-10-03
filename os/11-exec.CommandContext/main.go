package main

import (
	"bufio"
	"context"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(3))

	cmd := exec.CommandContext(ctx, "bash", "-c", "tcpdump -nni any -l")
	pipe, _ := cmd.StdoutPipe()

	go func() {
		buf := bufio.NewScanner(pipe)
		for buf.Scan() {
			log.Printf("%v\n", buf.Text())

		}
	}()

	err := cmd.Start()
	if err != nil {
		log.Printf("error: cmd.Start failed:%v\n", err)
	}

	cmd.Wait()
	cancel()

}
