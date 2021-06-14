package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type aWord struct {
	word      string
	timeStamp time.Time
}

func readWord(fh *os.File) (aWord string, n int) {
	myWord := []byte{}
	var err error
	for {
		buf := make([]byte, 1)
		n, err = fh.Read(buf)
		if err != nil {
			log.Println("error: reading file: ", err)
			break
		}

		myWord = append(myWord, buf...)

		if buf[0] == ' ' || buf[0] == '\n' {
			fmt.Println("found space")
			break
		}
	}
	return string(myWord), n
}

func main() {
	//Open a file, use 'test.txt' as default if none specified, or use filename given as input parameter.
	fileName := flag.String("fileName", "test.txt", "The filename to be checked")
	//Open file
	fh, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("error: opening file: ", err)
		os.Exit(1)
	}
	defer fh.Close()

	words := []aWord{}   //slice to hold all the words found
	wordFound := aWord{} //a single word found
	for {
		var n int
		wordFound.word, n = readWord(fh)
		wordFound.timeStamp = time.Now()
		words = append(words, wordFound)
		if n == 0 {
			break
		}
	}

	fmt.Println(words)

}
