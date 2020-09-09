package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	output, err := exec.Command("ifconfig").Output()
	if err != nil {
		log.Printf("error exec.Command failed: %v\n", err)
		os.Exit(1)
	}

	buf := bytes.NewBuffer(output)

	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, ": flags")
		if strings.HasPrefix(l[0], "en") {
			fmt.Printf("%#v\n", l[0])
		}
	}

}
