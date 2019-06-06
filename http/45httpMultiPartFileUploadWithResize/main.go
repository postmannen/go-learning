/*
	Test file upload on web page with MultiPart web page.
	Will read the whole file into memory,
	and write it all back to a temporary file.
*/

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nfnt/resize"
)

func shrinkImage(inReader io.Reader, outWriter io.Writer, size uint) error {
	decIm, _, err := image.Decode(inReader)
	if err != nil {
		return err
	}

	rezIm := resize.Resize(size, 0, decIm, resize.Lanczos3)

	err = jpeg.Encode(outWriter, rezIm, nil)
	if err != nil {
		return err
	}

	return nil
}

const (
	mainImageSize = 700
	thumbnailSize = 200
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

	// ------------------------- Creating main image ----------------------------------
	
	mainOutFile, err := ioutil.TempFile("./", "tmp100-*.jpg")
	if err != nil {
		log.Println("error: creating TempFile: ", err)
	}

	if err := shrinkImage(inFile, mainOutFile, thumbnailSize); err != nil {
		log.Println("error: shrink image failed: ", err)
	}

	mainOutFile.Close()
	// ------------------------- Creating thumbnail ----------------------------------
	thumbOutFile, err := ioutil.TempFile("./", "tmp400-*.jpg")
	if err != nil {
		log.Println("error: creating TempFile: ", err)
	}

	_, err = inFile.Seek(0, 0)
	if err != nil {
		log.Println("error: Failed seek to the start of read file: ", err)
	}

	if err := shrinkImage(inFile, thumbOutFile, mainImageSize); err != nil {
		log.Println("error: shrink image failed: ", err)
	}

	thumbOutFile.Close()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/upload", uploadFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("error: ListenAndServer failed: ", err)
	}
}
