package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

func fileOpen() error {
	fp, err := os.Open("a-not-existing-file.txt")
	if err != nil {
		return errors.Wrap(err, "Opening file problem")
	}
	defer fp.Close()
	return nil
}

func main() {
	err := fileOpen()
	if err != nil {
		//%v will give you the normal userspace error
		log.Printf("With percent v : %v\n", err)
		//%+v will give you the error, and the call stack
		log.Printf("With percent + v : %+v\n", err)
	}
}
