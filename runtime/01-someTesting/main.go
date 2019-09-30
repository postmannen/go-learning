package main

import (
	"fmt"
	"runtime"
)

func main() {

	s := struct {
		intOne int
		intTwo int
	}{
		intOne: 100,
		intTwo: 200,
	}

	fmt.Printf("unsafe size of s = %v\n\n", s)

	memstat := runtime.MemStats{}
	runtime.ReadMemStats(&memstat)
	// Alloc - currently allocated number of bytes on the heap,
	// TotalAlloc - cumulative max bytes allocated on the heap (will not decrease),
	// Sys - total memory obtained from the OS,
	// Mallocs and Frees - number of allocations, deallocations, and live objects (mallocs - frees),
	// PauseTotalNs - total GC pauses since the app has started,
	// NumGC - number of completed GC cycles
	fmt.Printf("memstats.Alloc = %v\n", memstat.Alloc)
	fmt.Printf("memstats.TotalAlloc = %v\n", memstat.TotalAlloc)
	fmt.Printf("memstats.Sys = %v\n", memstat.Sys)
	fmt.Printf("memstats.Mallocs = %v\n", memstat.Mallocs)
	fmt.Printf("memstats.PauseTotalNs = %v\n", memstat.PauseTotalNs)
	fmt.Printf("memstats.NumGC = %v\n", memstat.NumGC)

	fmt.Printf("\nNumber of go routines %+v\n", runtime.NumGoroutine())

}
