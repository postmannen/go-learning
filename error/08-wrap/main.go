package main

import (
	"errors"
	"fmt"
)

var errTest1 = errors.New("test1 error")

func returnErr() error {
	return fmt.Errorf("error: some error: %w", errTest1)
}

func main() {
	err := returnErr()
	if err != nil {
		e := errors.Unwrap(err)
		fmt.Printf("%v\n", e)
	}

	if errors.Is(err, errTest1) {
		fmt.Printf("is errTest1\n")
	}

}
