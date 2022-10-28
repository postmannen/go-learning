package main

import (
	"fmt"
	"strings"
)

func main() {
	sb := strings.Builder{}
	sb.WriteString("monkey")
	sb.WriteString(" man")

	fmt.Printf("%v\n", sb.String())
}
