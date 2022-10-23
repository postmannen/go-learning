package main

import (
	"log"
	"sync"
	"time"
)

type expireMap struct {
	mu       sync.Mutex
	m        map[int]string
	deleteCh chan chan int
}

func newExpireMap() *expireMap {
	e := expireMap{
		m:        make(map[int]string),
		deleteCh: make(chan chan int),
	}

	return &e
}

func (e *expireMap) start() {
	go func() {
		for k := range e.deleteCh {
			key := <-k
			e.mu.Lock()
			value := e.m[key]
			delete(e.m, key)
			e.mu.Unlock()
			log.Printf("info: delete cache value %v with key %v\n", value, key)
		}
	}()
}

func (e *expireMap) add(timeSeconds int, key int, value string) {
	go func() {
		f := newValue(e.deleteCh, timeSeconds, key)
		e.mu.Lock()
		go f()
		e.m[key] = value
		e.mu.Unlock()
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

func main() {
	em := newExpireMap()
	em.start()

	em.add(2, 1, "horse")
	em.add(4, 2, "sheep")
	em.add(1, 3, "pig")

	time.Sleep(time.Second * 5)

}
