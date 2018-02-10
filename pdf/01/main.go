package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

type billLine struct {
	description string
	price       float64
}

var lineHeight float64

func main() {
	billLines := []billLine{
		{description: "hest",
			price: 10},
		{description: "ku",
			price: 20},
		{description: "gris",
			price: 30},
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	fmt.Printf("pdf is of type %T\n", pdf)
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	lineHeight = 10

	for _, v := range billLines {
		pdf.Writef(lineHeight, "%-20v : %v\n", v.description, v.price)
		pdf.Ln(0)

	}

	//pdf.Cell(40, 10, "Hello, world")
	//pdf.Cell(70, 20, "test")
	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println("ERROR : ", err)
}
