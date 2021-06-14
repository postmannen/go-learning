package main

import (
	"fmt"
)

func main() {
	apekatt := "apekatt"
	hest := "hest"
	ku := "ku"

	//%15v will give the variable a print width of 15 spaces total
	//right aligned
	fmt.Printf("%15v %15v %15v\n", apekatt, hest, ku)
	fmt.Printf("%15v %15v %15v\n", hest, ku, apekatt)
	fmt.Printf("%15v %15v %15v\n", ku, apekatt, hest)

	fmt.Println()

	//%-15v will give the variable a print width of 15 spaces total
	//left aligned
	fmt.Printf("%-15v %-15v %-15v\n", apekatt, hest, ku)
	fmt.Printf("%-15v %-15v %-15v\n", hest, ku, apekatt)
	fmt.Printf("%-15v %-15v %-15v\n", ku, apekatt, hest)
}
