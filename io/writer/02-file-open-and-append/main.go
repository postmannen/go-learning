package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "Dette er en test med Fprint\n")
	io.WriteString(os.Stdout, "Dette er en test med io.WriteString\n")

	f, err := os.OpenFile("test.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	//fmt.Fprint(f, "apekatt") //writes apekatt at the beginning of the file, over existing text
	_, err = f.Write([]byte("Tekst som blir lagt til\n"))
	if err != nil {
		fmt.Println("Error, failed writing to file: ", err)
	}

	f.WriteString("tekst som blir lagt til med f.WriteString\n")

	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Error: ioutil.ReadAll : ", err)
	}

	fmt.Printf("%v ", string(content))

}
