package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data := []byte("This is some data that we want to compress")

	var buf bytes.Buffer

	{
		zw := gzip.NewWriter(&buf)
		n, err := zw.Write(data)
		if err != nil {
			log.Printf("error: gzip write failed: %v\n", err)
			return
		}
		log.Printf("info: n=%v bytes written\n", n)
		zw.Close()

		fmt.Printf("compressed data = %s\n", buf.Bytes())
	}

	{
		zr, err := gzip.NewReader(&buf)
		if err != nil {
			log.Printf("error: failed to create gzip reader: %v\n", err)
		}

		b, err := ioutil.ReadAll(zr)
		if err != nil {
			log.Printf("error: gzip read failed: %v\n", err)
			return
		}

		zr.Close()

		fmt.Printf("un-compressed data = %s\n", b)
	}
}
