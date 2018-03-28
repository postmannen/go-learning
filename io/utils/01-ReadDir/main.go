package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*reads the content of the directory, and returns whats found as type []os.FileInfo
		type FileInfo interface {
	        Name() string       // base name of the file
	        Size() int64        // length in bytes for regular files; system-dependent for others
	        Mode() FileMode     // file mode bits
	        ModTime() time.Time // modification time
	        IsDir() bool        // abbreviation for Mode().IsDir()
	        Sys() interface{}   // underlying data source (can return nil)
	}
	*/
	files, err := ioutil.ReadDir("/")
	if err != nil {
		fmt.Println("Error: Readdir: ", err)
	}

	for _, v := range files {
		if v.IsDir() {
			fmt.Println("Directory = ", v.Name())
		} else {
			fmt.Println("File = ", v.Name())
		}

	}
}
