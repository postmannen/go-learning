package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	fh, err := os.Open("bt.jpg")
	if err != nil {
		fmt.Println("error: os.Open: ", err)
	}
	defer fh.Close()

	decIm, _, err := image.Decode(fh)
	if err != nil {
		fmt.Println("error: image.Decode: ", err)
	}

	rezIm := resize.Resize(400, 0, decIm, resize.Lanczos3)
	oFh, err := os.Create("bt400.jpg")
	if err != nil {
		fmt.Println("error: os.Create: ", err)
	}
	defer oFh.Close()

	err = jpeg.Encode(oFh, rezIm, nil)
	if err != nil {
		fmt.Println("error: image.Encode: ", err)
	}
}
