package main

import (
	"fmt"
	"strings"
	"time"
)

func join[T fmt.Stringer](elements []T, separator string) string {
	sb := strings.Builder{}
	for i, v := range elements {
		if i > 0 {
			sb.WriteString(separator)
		}
		sb.WriteString(v.String())
	}

	return sb.String()
}

func main() {
	e := []time.Duration{time.Hour, time.Microsecond, time.Millisecond, time.Duration(time.Now().Day())}
	s := join(e, ", ")

	fmt.Printf("%v\n", s)

}
