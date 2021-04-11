package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Read the content of all the files in a directory, create
// a sha256 hash for the content of each file read, then put
// the hash as a key in a map with the value set to the file
// name.
// Check if a hash of the content are already present, and if
// it is, append the filename to the slice map value.
func main() {
	fi, err := ioutil.ReadDir("./")
	if err != nil {
		log.Printf("error: ReadDir: %v\n", err)
	}

	hashMap := map[string][]string{}

	for _, v := range fi {
		fh, err := os.Open(v.Name())
		if err != nil {
			log.Printf("error: Open: %v\n", err)
			continue
		}

		sha := sha256.New()
		_, err = io.Copy(sha, fh)
		if err != nil {
			log.Printf("error: io.Copy: %v\n", err)
			continue
		}

		shaSum := sha.Sum(nil)

		base64.StdEncoding.EncodeToString(shaSum)
		fmt.Printf("filename = %v, shaSum = %v\n", v.Name(), shaSum)

		// Check if the file hash is present in the map, and add it if it is not.
		mapV, ok := hashMap[string(shaSum)]
		if !ok {
			hashMap[string(shaSum)] = []string{v.Name()}
			continue
		}

		// If the file hash was already in the map, append the name of the current
		// file to the existing map value.
		mapV = append(mapV, v.Name())
		hashMap[string(shaSum)] = mapV
	}

	fmt.Printf("\n--- Content of map:\n")
	for _, v := range hashMap {
		fmt.Printf("%v\n", v)
	}
}
