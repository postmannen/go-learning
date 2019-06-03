/*
	Test file upload on web page with MultiPart web page.
	Will read the whole file into memory,
	and write it all back to a temporary file.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Entering uploadFile handler...")
	var err error

	//Takes max size of form to parse.
	err = r.ParseMultipartForm(10000000)
	if err != nil {
		log.Println("error: ParseMultipartForm: ", err)
	}

	//Get a handler for the file found in the web form.
	//Returns the first file for the provided key.
	fileWeb, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Println("err: failed to return web file: ", err)
	}
	defer fileWeb.Close()
	fmt.Printf("File uploaded : %v\n", fileHeader.Filename)
	fmt.Printf("File size : %v\n", fileHeader.Size)
	fmt.Printf("File MIME header : %v\n", fileHeader.Header)

	//Open a file for writing the uploaded file to.
	tempFile, err := ioutil.TempFile("./", "tmpfile-*.jpg")
	if err != nil {
		log.Println("error: creating TempFile: ", err)
	}
	defer tempFile.Close()

	//Read the whole file into a []byte
	fileBytes, err := ioutil.ReadAll(fileWeb)
	if err != nil {
		log.Println("error: Read whole web file: ", err)
	}

	//Write the file thats being read into the tmp file.
	n, err := tempFile.Write(fileBytes)
	if err != nil {
		log.Printf("error: Write file failed n=%v, err %v\n", n, err)
	}
	fmt.Println("Successfully upload file")

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/upload", uploadFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("error: ListenAndServer failed: ", err)
	}
}
