// Test with
// go run main.go -startCLISubscriber="a" -startCLISubscriber="b"

package main

import (
	"flag"
	"fmt"
)

type StartCLISubscriber struct {
	value []string
}

func (f *StartCLISubscriber) String() string {
	return ""
}

func (f *StartCLISubscriber) Set(s string) error {
	f.value = append(f.value, s)
	return nil
}

func main() {
	var startCLISubscriber StartCLISubscriber

	flag.Var(&startCLISubscriber, "startCLISubscriber", "enter value")
	flag.Parse()

	fmt.Printf("%#v\n", startCLISubscriber)
}
