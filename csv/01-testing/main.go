package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
)

func main() {

	raw := []byte("a,b,c,1,2,3")
	rd := bytes.NewReader(raw)

	csvRd := csv.NewReader(rd)

	s, err := csvRd.Read()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(s)

}
