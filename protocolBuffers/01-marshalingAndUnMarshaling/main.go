package main

import "github.com/golang/protobuf/proto"

import "log"

import "fmt"

func main() {
	bt := Person{
		Name:    "Bjørn Tore",
		HouseNR: 945,
	}

	b, err := proto.Marshal(&bt)
	if err != nil {
		log.Printf("marshaling failed: %v\n", err)
	}

	fmt.Println(b)

	bt2 := Person{}
	proto.Unmarshal(b, &bt2)

	fmt.Printf("Unmarshaled Person = %+v\n", bt2)
	fmt.Println(bt2.Name)
	fmt.Println(bt2.HouseNR)
	fmt.Println(bt2.GetName())
	fmt.Println(bt2.GetHouseNR())

}
