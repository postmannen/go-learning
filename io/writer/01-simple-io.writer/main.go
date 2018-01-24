package main

import (
	"bytes"
	"fmt"
	"os"
)

const checkMark = "\u2713"
const xMark = "\u2717"

func main() {
	b1 := bytes.Buffer{}
	b1.Write([]byte("apekatt"))
	b1.WriteTo(os.Stdout)
	fmt.Println("\n------------")
	b1.WriteTo(os.Stdout) //buffer is empty ?
	fmt.Println("\n------------", checkMark, xMark)

}
