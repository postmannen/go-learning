package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type Foo struct {
	count int
	text  string
}

var before uint64
var first uint64
var second uint64

func main() {

	// Force GC to clear up
	runtime.GC()

	// Below is an example of using our PrintMemUsage() function
	// Print our starting memory usage (should be around 0mb)
	before = PrintMemUsage()

	a := &Foo{}
	first = PrintMemUsage()

	// Force GC to clear up
	runtime.GC()

	a = (*Foo)(nil)
	second = PrintMemUsage()

	fmt.Printf("Total: %v bytes\n", before)

	fmt.Printf("&Foo{}\t\t Total: %v bytes, allocated %v bytes, pointer %v bytes\n", first, first-before, unsafe.Sizeof(a))

	fmt.Printf("(*Foo)(nil)\t Total: %v bytes, allocated %v bytes, pointer %v bytes\n", second, second-before, unsafe.Sizeof(a))

	fmt.Println(a)
}

//PrintMemUsage prints memory allocations.
func PrintMemUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m.Alloc
}
