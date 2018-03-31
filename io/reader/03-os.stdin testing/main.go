package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	readBytes := []byte{}

	reader := os.Stdin
	p := make([]byte, 1)
	for {
		n, err := reader.Read(p)
		if err == io.EOF || p[0] == '\n' {
			break
		}
		readBytes = append(readBytes, p[:n]...)
		fmt.Println("chunc of bytes read = ", string(p[:n]))

	}

	fmt.Println("Complete string read = ", string(readBytes))

	cmd := exec.Command("ls -l")
	//cmd.Run()
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("command output error : ", err)
	}

	fmt.Println(string(out))
}
