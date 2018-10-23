package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Println("Failed to open file : ", err)
	}

	fIO := bufio.NewReader(f)

	lines := []string{}
	for {

		line, isPrefix, err := fIO.ReadLine()
		if err != nil {
			log.Println("Error reading line : ", err)
			break
		}
		lines = append(lines, string(line))
		fmt.Printf("Line = %v, isPrefix = %v \n", line, isPrefix)
	}

	fmt.Println(lines)
}
