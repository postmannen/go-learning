package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	out, err := cmd.Output()
	if err != nil {
		log.Println("error: cmd failed: ", err)
	}

	fmt.Println("out = ", string(out))
}
