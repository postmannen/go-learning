package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type myData struct {
	word      string
	timeStamp time.Time
}

func (m myData) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//open file
	fh, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("error: opening file: ", err)
		os.Exit(1)
	}

	allText, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println("error: reading file: ", err)
	}

	tmpSplit := strings.Split(string(allText), " ")

	words := []myData{}
	tmpWord := myData{}
	for _, v := range tmpSplit {
		tmpWord.word = v
		tmpWord.timeStamp = time.Now()
		words = append(words, tmpWord)
	}

	fmt.Println(words)
	//http.Handle("/", words)
	http.ListenAndServe(":8080", nil)
}
