package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(40, 10, "Hello, world")
	pdf.Cell(70, 20, "test")
	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println("ERROR : ", err)
}
