package arguments

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
)

/*
u8 1 unsigned 8bit value
i8 1 signed 8bit value
u16 2 unsigned 16bit value
i16 2 signed 16bit value
u32 4 unsigned 32bit value
i32 4 signed 32bit value
u64 8 unsigned 64bit value
i64 8 signed 64bit value
float 4 IEEE-754 single precision
double 8 IEEE-754 double precision
string * Null terminated string (C-String)
(Variable size)
enum 4 Per command defined enum
*/

// ------------------------------------------------------------------------------------

// uint8Type makes a type for uint8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type uint8Type struct {
	length int
}

// U8 is the u8 type
var U8 = uint8Type{
	length: 1,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *uint8Type) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *uint8Type) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *uint8Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the uint8Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val uint8

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// int8Type makes a type for int8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type int8Type struct {
	length int
}

// I8 is the i8 type
var I8 = int8Type{
	length: 1,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *int8Type) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *int8Type) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *int8Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the int8Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val int8

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// uint16Type makes a type for int8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type uint16Type struct {
	length int
}

// U16 is the u16 type
var U16 = uint16Type{
	length: 2,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *uint16Type) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *uint16Type) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *uint16Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the uint16Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val uint16

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// int16Type makes a type for int8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type int16Type struct {
	length int
}

// I16 is the i16 type
var I16 = uint16Type{
	length: 2,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *int16Type) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *int16Type) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *int16Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the int16Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val int16

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// uint32Type
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type uint32Type struct {
	length int
}

// U32 is the u32 type.
var U32 = uint32Type{
	length: 4,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *uint32Type) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *uint32Type) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *uint32Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the uint32Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val uint32

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// int32Type
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type int32Type struct {
	length int
}

// I32 is the i32 type
var I32 = int32Type{
	length: 4,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *int32Type) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *int32Type) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *int32Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the int32Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val int32

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// uint64Type
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type uint64Type struct {
	length int
}

// U64 is the u64 type
var U64 = uint64Type{
	length: 8,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *uint64Type) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *uint64Type) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *uint64Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the uint64Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val uint64

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// int64Type
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type int64Type struct {
	length int
}

// I64 is the i64 type
var I64 = int64Type{
	length: 8,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *int64Type) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *int64Type) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *int64Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the int64Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val int64

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// float32Type makes a type for int8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type float32Type struct {
	length int
}

// Float makes a type for float32 data, and tells the length of bytes it
// is made of.
var Float = float32Type{
	length: 4,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *float32Type) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *float32Type) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *float32Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the float32.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val float32

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// float64Type
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type float64Type struct {
	length int
}

// Double makes a type for float64 data, and tells the length of bytes it
// is made of.
var Double = float64Type{
	length: 8,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *float64Type) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *float64Type) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *float64Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the float64.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val float64

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// stringType
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type stringType struct {
	length int
}

// Stringx is the string type
var Stringx = stringType{
	length: 0,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *stringType) getLength() int {
	return f.length
}

// setLength will set the value of the length field t.length for type.
func (f *stringType) setLength(length int) {
	f.length = length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *stringType) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the stringx.argDecode method, b = %v\n", b)

	fmt.Printf("Content = %#v\n", string(b))

	return string(b), nil

	//TODO: Implement string argDecode logic !!!

}

// ------------------------------------------------------------------------------------

// enumType
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
// Since enum is 4 bytes, we use an int32
type enumType struct {
	length int
}

// Enum is the enum type
var Enum = enumType{
	length: 4,
}

// setLength will set the value of the length field t.length for type.
func (f *enumType) setLength(length int) {
	f.length = length
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *enumType) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *enumType) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the enum.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val int32

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	fmt.Printf("Content = %#v\n", val)

	return val, nil
}

// ------------------------------------------------------------------------------------

// ------------------------------------------------------------------------------------

// InsertArgValueIntoStruct takes the struct to fill as a pointer value,
// and the arguments as a slice of []interface{} as input.
// It will use reflect to loop over the struct fields, and set the correct
// value for each field from the []argvalues.
// NB: The order of the []argvalues have to be the same as the order of the
// elements in the struct.
func InsertArgValueIntoStruct(argStruct interface{}, argValues []interface{}) error {
	dataValue := reflect.ValueOf(argStruct)
	if dataValue.Kind() != reflect.Ptr {
		panic("not a pointer")
	}

	dataElements := dataValue.Elem()

	// Check if the number of fields in the slice and the number of argValues
	// match, return with an error.
	if dataElements.NumField() != len(argValues) {
		err := fmt.Errorf(
			"Number of fields for argStruct and argValues are not the same, check length of both variables given as input to insertArgValueIntoStruct")

		return err
	}

	//this loops through the fields
	for i := 0; i < dataElements.NumField(); i++ { // iterates through every struct type field
		//k := elements.Kind()
		dataType := dataElements.Type().Field(i).Type // returns the tag string
		dataField := dataElements.Field(i)            // returns the content of the struct type field

		argVal := reflect.ValueOf(argValues[i])
		fmt.Printf("argVal = %+v, type = %T\n", argVal, argVal)

		// Check what the types it is, and when the correct type for the field is
		// found, insert the value into the struct field.
		switch dataType.String() {
		case "int":
			fmt.Printf("Reflecting INT\n")
			v := argVal.Int()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetInt(v)
		case "uint8":
			fmt.Printf("Reflecting uint8\n")
			v := argVal.Uint()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetUint(v)
		case "int8":
			fmt.Printf("Reflecting INT8\n")
			v := argVal.Int()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetInt(v)
		case "uint16":
			fmt.Printf("Reflecting uint16\n")
			v := argVal.Uint()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetUint(v)
		case "int16":
			fmt.Printf("Reflecting INT16\n")
			v := argVal.Int()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetInt(v)
		case "uint32":
			fmt.Printf("Reflecting uint32\n")
			v := argVal.Uint()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetUint(v)
		case "int32":
			fmt.Printf("Reflecting INT32\n")
			v := argVal.Int()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetInt(v)
		case "uint64":
			fmt.Printf("Reflecting uint64\n")
			v := argVal.Uint()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetUint(v)
		case "int64":
			fmt.Printf("Reflecting INT64\n")
			v := argVal.Int()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetInt(v)
		case "float32":
			fmt.Printf("Reflecting float32\n")
			v := argVal.Float()
			dataField.SetFloat(v)
		case "float64":
			fmt.Printf("Reflecting float64\n")
			v := argVal.Float()
			dataField.SetFloat(v)
		case "string":
			fmt.Printf("Reflecting string\n")
			v := argVal.String()
			dataField.SetString(v)
		}
	}
	return nil
}

// ------------------------------------------------------------------------------------
// argDecoder is an interface type which tells that any type that
// have the methods argDecode([]byte) error, and getLength() int
// is of the interface type argDecoder.
type argDecoder interface {
	argDecode([]byte) (interface{}, error)
	getLength() int
	setLength(int)
}

func getLengthOfData(b []byte) (int, error) {
	// Figure out the length of the string
	for i := 0; i < cap(b); i++ {
		//fmt.Printf("%+v, of type %T\n", b[i], b[i])

		//fmt.Println("i = ", i)
		if b[i] == 0 {
			//fmt.Println("lengthString = ", i)

			// add 1 to jump to the 0
			return i + 1, nil
		}

	}

	err := fmt.Errorf("no string bytes found, returning 0")
	return 0, err
}

// Decoder is a type for keeping track of the start position of the
// data to read in a slice.
type Decoder struct {
	position int
}

// NewDecoder will return a new argument decoder type.
func NewDecoder() *Decoder {
	return &Decoder{}
}

// DecodeArgs takes a []byte and any number of the interface type *argDecoder
// is input.
// It will loop through the argDecoder methods, and run the concrete types method,
// one by one until all methods are done.
// The method will use the getLength() method to know the size of the portion of data
// to work with, and increase the the position with the size of the last data read to
// know where the next piece of data starts.
// TODO: Make logic check if there are given the correct amount of argDecoders to
// handle the length of the data slice given as input, and return error if they don't
// match.
func (as *Decoder) DecodeArgs(argStruct interface{}, d []byte, a ...argDecoder) ([]interface{}, error) {
	as.position = 0
	argumentSlice := []interface{}{}
	for _, v := range a {
		fmt.Println("------------------Decoding byte or bytes----------------------")

		// The string type has initial length to 0. The reason is that we never know
		// the length of a string since it ends with a value of 0 in the slice.
		// We need to check if length = 0, and update the length for the type.
		var length int
		if v.getLength() == 0 {
			l, err := getLengthOfData(d[as.position:])
			if err != nil {
				log.Println("error: argumentsToDecode: failed to get the length :", err)
			}
			length = l
			//v.setLength(l)
			fmt.Println("The value after setLength = ", length)
		} else {
			length = v.getLength()
		}

		val, err := v.argDecode(d[as.position : as.position+length])
		if err != nil {
			return nil, err
		}

		// Putting the values into a slice, to iterate later.
		argumentSlice = append(argumentSlice, val)
		fmt.Printf("val = %+v, type = %T\n", val, val)

		l := length
		as.position += l
		fmt.Println("--------------------------------------------------------------")
	}

	return argumentSlice, nil
}
