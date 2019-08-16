package main

import (
	"fmt"
	"reflect"
)

type post struct {
	ArgumentHeight int     `type:"int"`
	ArgumenAngle   float32 `type:"float"`
}

func checkSwitch(data interface{}, argValues []interface{}) {
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Ptr {
		panic("not a pointer")
	}

	dataElements := dataValue.Elem()

	//this loops through the fields
	for i := 0; i < dataElements.NumField(); i++ { // iterates through every struct type field
		//k := elements.Kind()
		dataType := dataElements.Type().Field(i).Type // returns the tag string
		dataField := dataElements.Field(i)            // returns the content of the struct type field

		argVal := reflect.ValueOf(argValues[i])
		fmt.Printf("argVal = %+v, type = %T\n", argVal, argVal)

		switch dataType.String() {
		case "int":
			v := argVal.Int()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetInt(v)
		case "float64":
			v := argVal.Float()
			dataField.SetFloat(v)
		case "float32":
			v := argVal.Float()
			dataField.SetFloat(v)
		}
	}
}

func main() {
	p := &post{ArgumentHeight: 10, ArgumenAngle: 10.5}
	// testing
	s := []interface{}{int(3333), float32(3333.3)}

	fmt.Println(p)
	checkSwitch(p, s)
	fmt.Println(p)
}

