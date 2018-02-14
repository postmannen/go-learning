package main

import (
	"fmt"
	"strings"
)

func main() {

	myReader := strings.NewReader("Dette er en liten test som er mekket sammen")
	fmt.Println("-----------------------------------")
	fmt.Printf("myReader type = %T, og inneholdet er = %v\n", myReader, myReader)

	//lag en slice of bytes som skal fungere som transfer buffer
	//"len" størelsen på slice som blir initiert, bestemmer hvor stort transfer bufferet skal være
	mySliceOfByte := make([]byte, 16)

	for {
		n, err := myReader.Read(mySliceOfByte)
		fmt.Println("n = ", n)
		fmt.Println("err = ", err)
		if err != nil {
			break
		}
	}
	fmt.Println(mySliceOfByte)

}
