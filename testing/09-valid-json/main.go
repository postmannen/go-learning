package main

import (
	"encoding/json"
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

	var failure int

	for _, f := range jsonFilepaths {
		func() {
			fh, err := os.Open(f)
			if err != nil {
				log.Printf("error: failed to open file for reading: %v\n", err)
				os.Exit(1)
			}
			defer fh.Close()

			js, err := io.ReadAll(fh)
			if err != nil {
				log.Printf("error: failed to read file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf(" * Content of file %v : %v\n", f, string(js))

			var data map[string]any

			if len(js) < 1 {
				fmt.Printf(" * the file %v was empty: ", f)
				return
			}

			err = json.Unmarshal(js, &data)
			if err != nil {
				log.Printf("error: for file %v, failed to unmarshal data: %v\n", f, err)
				failure = 1
			}
		}()

	}

	if failure != 0 {
		os.Exit(failure)
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
