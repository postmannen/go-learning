package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

func main() {
	// Open the original image.
	fp, err := os.Open("./bt.jpg")
	if err != nil {
		fmt.Println("error: os.Open: ", err)
	}
	// Put the closing of the opened filed as defered, so it will
	// be closed when main() exits.
	defer fp.Close()

	//DecodeConfig decodes all the config of the image.
	imConfig, imType, err := image.DecodeConfig(fp)
	if err != nil {
		fmt.Println("error: image.DecodeConfig: ", err)
	}
	fmt.Printf("imConfig = %v, imType = %v \n", imConfig, imType)

	//seek will return back to starting position.
	ret, err := fp.Seek(0, 0)
	if err != nil {
		fmt.Println("error: fp.Seek: ", err)
	}
	fmt.Println("ret = ", ret)

	//Decode will decode into a raw format.
	im, imS, err := image.Decode(fp)
	if err != nil {
		fmt.Println("error: image.Decode: ", err)
	}
	fmt.Printf("imS=%v\n", imS)

	//Create a file for writing, convert to png and write it as a png file.
	of, err := os.Create("bt.png")
	if err != nil {
		fmt.Println("error: os.Create: ", err)
	}
	defer of.Close()

	//Encode the raw image encoded earlier into png format.
	err = png.Encode(of, im)
	if err != nil {
		fmt.Println("error: image.Encode: ", err)
	}

}
