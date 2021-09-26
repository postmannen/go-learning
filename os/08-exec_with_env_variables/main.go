package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	fh, err := os.Open("test.txt")
	if err != nil {
		log.Printf("error: failed to open file: %v\n", err)
		return
	}
	defer fh.Close()

	b, err := io.ReadAll(fh)
	if err != nil {
		log.Printf("error: ReadAll failed: %v\n", err)
		return
	}

	data := fmt.Sprintf("HORSE=%s", b)

	cmd := exec.Command("bash", "-c", "echo $HORSE> horse.txt")
	cmd.Env = append(cmd.Env, data)
	err = cmd.Run()
	if err != nil {
		log.Printf("error: exec: %v\n", err)
		return
	}
}
