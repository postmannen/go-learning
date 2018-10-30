package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "ardrone3.xml"
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("Error: os.Open: %v\n", err)
	}
	defer f.Close()

	//bufio lets us read files line by line
	fReader := bufio.NewReader(f)

	lineNR := 1
	for {
		//fmt.Println("----------Line nr = ", lineNR)
		//read a line
		line, _, err := fReader.ReadLine()
		if err != nil {
			log.Printf("Error: bufio.ReadLine: %v\n", err)
			break
		}

		//Remove leading spaces in line
		tmpLine := strings.TrimSpace(string(line))
		line = []byte(tmpLine)
		printLine(line)

		err = checkAngleBracket(line)
		if err != nil {
			fmt.Println(err)
		}

		lineNR++
	}

}

func printLine(line []byte) {
	//fmt.Printf("Line : %v \n Type %T\n", line, line)
	for i := 0; i < len(line); i++ {
		character := string(line[i])
		fmt.Print(character)

	}
	fmt.Println()
}

func checkAngleBracket(line []byte) error {
	//Check for opening and closing angle bracket.
	//Will return nil if both start and end bracker was found.
	for i := 0; i < len(line); i++ {
		character := string(line[i])
		if character == "<" {
			ii := 0
			for {
				if string(line[ii]) == ">" {
					fmt.Println("Found closing angular bracket at position: ", ii)
					break
				}
				if ii == len(line)-1 {
					fmt.Println("Missing ending angular bracket at line: ")
					return errors.New("Missing ending angular bracket at line: ")
				}
				ii++
			}
		}
	}
	return nil
}
