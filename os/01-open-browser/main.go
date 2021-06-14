package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)

	switch runtime.GOOS {
	case "darwin":
		fmt.Println("The OS which is chosen is MacOs")
		cmd := exec.Command("open", "https://google.com")
		cmd.Run()

	}
}
