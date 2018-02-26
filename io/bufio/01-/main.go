package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileHandle, err := os.Open("file.txt")
	if err != nil {
		log.Println("Error: opening file : ", err)
	}

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
