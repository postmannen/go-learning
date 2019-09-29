package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type someStruct struct {
	count int
	text  string
}

var before uint64
var first uint64
var second uint64

func main() {

	// Force GC to clear up
	runtime.GC()

	before = getTotalAllocations()
	fmt.Printf("Total before creating any variables: %v bytes\n\n", before)

	a := &someStruct{}
	first = getTotalAllocations()
	fmt.Printf("&someStruct{}\t\t Total: %v bytes, allocated %v bytes, pointer %v bytes\n\n", first, first-before, unsafe.Sizeof(a))

	// Force GC to clear up
	runtime.GC()

	a = (*someStruct)(nil)
	second = getTotalAllocations()
	fmt.Printf("(*someStruct)(nil)\t Total: %v bytes, allocated %v bytes, pointer %v bytes\n\n", second, second-before, unsafe.Sizeof(a))

	fmt.Println(a)
}

//getTotalAllocations prints memory allocations.
func getTotalAllocations() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m.Alloc
}
