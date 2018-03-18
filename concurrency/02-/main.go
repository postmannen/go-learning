package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func main() {
	var counter int //defaults to nil, which is 0 for int

	/*
		Without Mutex there is no guarantee what order the read and write operations will happen.
		The go routines will be distributed over different threads,
		and might do so the the order executed ain't happening in the order from 1->20
	*/

	for i := 0; i <= 20; i++ {
		go incWithoutMutex(&counter)
	}
	time.Sleep(1 * time.Second)

	/*
		With mutex Lock(), it is made certain that only one thread is writing to the variable at a time,
	*/
	counter = 0

	for i := 0; i <= 20; i++ {
		go incWithMutex(&counter)
	}
	time.Sleep(1 * time.Second)

}

func incWithoutMutex(n *int) {
	*n++
	fmt.Println("Counter = ", *n)
}

func incWithMutex(n *int) {
	lock.Lock()
	defer lock.Unlock()
	*n++
	fmt.Println("Counter = ", *n)
}
