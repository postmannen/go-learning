package main

import (
	"fmt"
	"sync"
)

var x int

func incX(ch chan int) {
	x = x + 1
	ch <- x
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	ch := make(chan int)

	count := 1000

	wg.Add(count)

	go func() {
		for i := 0; i <= count; i++ {
			go incX(ch)
		}
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

}
