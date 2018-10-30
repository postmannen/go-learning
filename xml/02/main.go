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
		//printLine(line)

		if tp := checkTagProject(line); tp.foundStart {
			fmt.Println("AAA found project start on lineNR : ", lineNR)
			if err := checkForClosingBracket(line); err != nil {
				log.Fatal("Error: missed closing bracket")
			}
		}

		if tp := checkTagProject(line); tp.foundEnd {
			fmt.Println("AAA found project end on lineNR : ", lineNR)
			if err := checkForClosingBracket(line); err != nil {
				log.Fatal("Error: missed closing bracket")
			}
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

type tagProject struct {
	foundStart bool
	foundEnd   bool
	name       string
	id         string
}

func checkTagProject(line []byte) tagProject {
	var tag string
	if len(line) > 0 {
		tag = string(line[0:8])
		if tag == "<project" {
			tp := tagProject{foundStart: true}

			//find word in []byte
			myWordString := "name="
			myWordByte := []byte(myWordString)
			foundWord := false

			for linePosition := 0; linePosition < len(line)-len(myWordByte); linePosition++ {
				wordPosition := 0
				for {
					if wordPosition >= len(myWordByte) {
						fmt.Println("Reached the end of the word, breaking out of word loop", linePosition, wordPosition)
						foundWord = true
						break
					}

					fmt.Println("Comparing : ", line[linePosition+wordPosition], myWordByte[wordPosition])

					if line[linePosition+wordPosition] == myWordByte[wordPosition] {
						fmt.Println("found letter: ", string(myWordByte[wordPosition]))
					} else {
						fmt.Println("Letters don't match, breaking out of inner loop")
						break
					}

					//if wordPosition == len(myWordByte)-1 {
					//	fmt.Println("Found the word")
					//	break
					//}

					wordPosition++
				}

				if foundWord {
					fmt.Println("Breaking out of outer loop")
					break
				}
			}

			return tp
		}

		tag = string(line[0:9])
		if tag == "</project" {
			tp := tagProject{foundEnd: true}
			return tp
		}

	}

	tp := tagProject{}
	return tp
}

func checkForClosingBracket(line []byte) error {
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

					return errors.New("Missing ending angular bracket")
				}
				ii++
			}
		}
	}
	return nil
}
