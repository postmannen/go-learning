package main

import (
	"io"
	"os"
)

/*
Bruker os for å åpne fil
Bruker bufio for å lese igjennom fila
*/

func main() {
	f, _ := os.Open("data.txt") //fil blir peker til *os.file , os brukes til å åpne fil
	//	input := bufio.NewScanner(f) //NewScanner tar io.Reader som input, og gir *Scanner som return

	//	for input.Scan() {
	//		fmt.Print(input.Text())
	//	}
	//fmt.Printf("%T\n", f)

	io.Copy(os.Stdout, f)

}
