//testing file reading in one go routine and sending whats read to other goroutine
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	ch1 := make(chan byte)

	dirContent, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println("error: ReadDir: ", err)
		os.Exit(1)
	}

	wFile, err := os.Create("/tmp/output.txt")
	if err != nil {
		fmt.Println("error: creating output file: ", err)
	}

	defer wFile.Close()

	//range all the files found in the directory
	for _, v := range dirContent {
		//open file
		rFile, err := os.Open(v.Name())
		if err != nil {
			fmt.Println("error: file open: ", err)
		}

		wg.Add(2)
		go read(ch1, rFile)
		go write(ch1, wFile)

	}

	wg.Wait()
}

func write(c1 chan byte, wf *os.File) {
	vv := []byte{0}
	for {
		//read from channel, ok variable will become false if channel is closed
		v, ok := <-c1
		if ok {
			vv[0] = v
			wf.Write(vv)
			fmt.Print(string(v))
		} else {
			wg.Done()
			break
		}
	}
}

//Will read 1 byte at a time of the file given as input,
//send that byte on the channel
//close the channel and exit when done
func read(c1 chan byte, fh *os.File) {
	//creating a slice of byte who can hold only 1 byte
	data := make([]byte, 1)
	for {
		_, err := fh.Read(data)
		//EOF or other error ?, close channel and exit for loop.
		if err != nil {
			close(c1)
			wg.Done()
			break
		}

		//since the data buffer is only 1 byte, we read only the first index of the slice
		c1 <- data[0]

		//fmt.Print(string(data))
	}
}
