package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Example of reading the content of one file, and then after each
// line read, write that line into a new file.
func main() {
	inFile, err := os.Open("bt-mini-2.jpg")
	if err != nil {
		log.Println("error: failed opening file: ", err)
	}
	defer inFile.Close()

	outFile, err := os.Create("outfile.jpg")
	if err != nil {
		log.Println("error: failed creating file: ", err)
	}
	defer outFile.Close()

	inReader := bufio.NewReader(inFile)
	outWriter := bufio.NewWriter(outFile)

	// Read bytes up until the occurence of a lineshift, and then
	// write the data read out to the output file.
	for {
		bs, rerr := inReader.ReadBytes('\n')

		nn, err := outWriter.Write(bs)
		if err != nil {
			log.Println("error: failed write: ", err)
		}
		fmt.Println(nn)

		if rerr != nil {
			log.Println("error: readBytes: ", err)
			break
		}
	}

	if err := outWriter.Flush(); err != nil {
		log.Println("error: flush failed: ", err)
	}
}
