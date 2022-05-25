package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	jsonFilepaths := getJsonFilePaths()
	fmt.Printf("file=%v, and the type is %T\n", jsonFilepaths, jsonFilepaths)

	for _, f := range jsonFilepaths {
		fh, err := os.Open(f)
		if err != nil {
			log.Printf("error: failed to open file for reading: %v\n", err)
			os.Exit(1)
		}
		// Remember to close the file!

		js, err := io.ReadAll(fh)
		if err != nil {
			log.Printf("error: failed to read file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf(" * Content of file %v : %v\n", f, string(js))

	}

}

func getJsonFilePaths() []string {
	files := []string{}

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(path, ".json") {
				fmt.Println(path)
				files = append(files, path)
			}

			return nil

		})
	if err != nil {
		log.Println(err)
	}

	return files
}
