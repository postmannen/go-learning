package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

type aWord struct {
	word      string
	timeStamp time.Time
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

	//define a new scanner
	scanner := bufio.NewScanner(fh)
	//set the split function to look for words
	scanner.Split(bufio.ScanWords)

	words := []string{}
	//do the scanning
	for scanner.Scan() {
		//return whats scanned as type string, and add it to the slice
		words = append(words, scanner.Text())
	}

	fmt.Println(words)
	fmt.Println("Length of words slice = ", len(words))
}
