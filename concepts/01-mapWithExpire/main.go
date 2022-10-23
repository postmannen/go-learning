package main

import (
	"log"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	deleteCh := make(chan chan int)

	m := make(map[int]string)

	go func() {
		for k := range deleteCh {
			key := <-k
			mu.Lock()
			value := m[key]
			delete(m, key)
			mu.Unlock()
			log.Printf("info: delete cache value %v with key %v\n", value, key)
		}
	}()

	addMapKeyValue(deleteCh, m, 2, 1, "horse")
	addMapKeyValue(deleteCh, m, 4, 2, "sheep")
	addMapKeyValue(deleteCh, m, 1, 3, "pig")

	time.Sleep(time.Second * 5)

}

func addMapKeyValue(deleteCh chan chan int, m map[int]string, timeSeconds int, key int, value string) {
	go func() {
		f := newValue(deleteCh, timeSeconds, key)
		mu.Lock()
		go f()
		m[key] = value
		mu.Unlock()
	}()
}

func newValue(deleteCh chan chan int, timeSeconds int, key int) func() {
	f := func() {
		key := key
		ch := make(chan int, 1)
		ch <- key

		ticker := time.NewTicker(time.Second * time.Duration(timeSeconds))
		// fmt.Printf(" * new.Ticker\n")
		for range ticker.C {
			// fmt.Printf(" * got ticker\n")
			deleteCh <- ch
			// fmt.Printf(" * put ch on deleteCh\n")
			return
		}
	}

	return f
}
