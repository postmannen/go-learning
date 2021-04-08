package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fp := filepath.Join("./a", "/b/c/", "d.doc")
	fmt.Printf("%v\n", fp)
	err := os.MkdirAll(fp, 0755)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
