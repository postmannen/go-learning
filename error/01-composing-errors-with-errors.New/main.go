package main

import (
	"errors"
	"log"
	"os"
)

func fileOpen() error {
	fp, err := os.Open("a-not-existing-file.txt")
	if err != nil {
		return errors.New("There was a problem opening the file : " + err.Error())
	}
	defer fp.Close()
	return nil
}

func main() {
	err := fileOpen()
	if err != nil {
		log.Println(err)
	}
}
