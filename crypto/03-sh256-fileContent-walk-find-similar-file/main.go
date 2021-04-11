package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Read the content of all the files in a directory, create
// a sha256 hash for the content of each file read, then put
// the hash as a key in a map with the value set to the file
// name.
// Check if a hash of the content are already present, and if
// it is, append the filename to the slice map value.
func main() {
	hashMap := map[string][]string{}

	err := filepath.Walk("../../", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		fh, err := os.Open(path)
		if err != nil {
			log.Printf("error: Open: info = %v,  %v\n", info, err)
			return err
		}

		sha := sha256.New()
		_, err = io.Copy(sha, fh)
		if err != nil {
			log.Printf("error: io.Copy: %v\n", err)
			return err
		}

		shaSum := sha.Sum(nil)

		base64.StdEncoding.EncodeToString(shaSum)
		//fmt.Printf("filename = %v, shaSum = %v\n", info.Name(), shaSum)

		// Check if the file hash is present in the map, and add it if it is not.
		mapV, ok := hashMap[string(shaSum)]
		if !ok {
			hashMap[string(shaSum)] = []string{path}
			return err
		}

		// If the file hash was already in the map, append the name of the current
		// file to the existing map value.
		mapV = append(mapV, path)
		hashMap[string(shaSum)] = mapV

		return nil
	})

	if err != nil {
		log.Printf("error: Walk failed: %v\n", err)
		return
	}

	fmt.Printf("\n--- Content of map:\n")
	for _, v := range hashMap {
		if len(v) > 1 {
			fmt.Printf("value = %v\n", v)
		}
	}
}
