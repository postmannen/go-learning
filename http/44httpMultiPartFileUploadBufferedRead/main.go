/*
	Test file upload on web page with MultiPart web page.
	Will read the whole file into memory,
	and write it all back to a temporary file.
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	var err error

	//Takes max size of form to parse.
	err = r.ParseMultipartForm(10000000)
	if err != nil {
		log.Println("error: ParseMultipartForm: ", err)
	}

	//Get a handler for the file found in the web form.
	//Returns the first file for the provided key.
	inFile, inFileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Println("err: failed to return web file: ", err)
	}
	defer inFile.Close()
	fmt.Printf("File uploaded : %v\n", inFileHeader.Filename)
	fmt.Printf("File size : %v\n", inFileHeader.Size)
	fmt.Printf("File MIME header : %v\n", inFileHeader.Header)

	//Open a file for writing the uploaded file to.
	tempFile, err := ioutil.TempFile("./", "tmpfile-*.jpg")
	if err != nil {
		log.Println("error: creating TempFile: ", err)
	}
	defer tempFile.Close()

	inScanner := bufio.NewScanner(inFile)
	//outWriter := bufio.NewWriter(tempFile)

	for {
		v := inScanner.Scan()
		if err := inScanner.Err(); err != nil {
			log.Println("error: scanner: ", err)
		}

		fmt.Println(len(inScanner.Bytes()))

		//_, err := outWriter.Write(inScanner.Bytes())
		//if err != nil {
		//	log.Println("error: write failed : ", err)
		//}
		//outWriter.Flush()

		if !v {
			log.Println("done scanning")
			break
		}
	}

	//http.Redirect(w, r, r.Header.Get("Referer"), 302)

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/upload", uploadFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("error: ListenAndServer failed: ", err)
	}
}
