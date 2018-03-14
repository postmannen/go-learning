package main

import (
	"fmt"
	"os"
)

func printToSTDOUT(a ...interface{}) {
	fmt.Fprintln(os.Stdout, a)
}

func main() {
	printToSTDOUT("apekatt", "hest")

}
