package main

import (
	"context"
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

func (e *expireMap) start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			select {
			case k := <-e.deleteCh:
				go func(k chan int) {
					key := <-k
					e.mu.Lock()
					value := e.m[key]
					delete(e.m, key)
					e.mu.Unlock()
					log.Printf("info: delete cache value %v with key %v, len of map %v\n", value, key, len(e.m))
				}(k)
			case <-ctx.Done():
				log.Printf("info: ctx.Done in start()\n")
				wg.Done()
				return
			}
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
		defer ticker.Stop()
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var wg sync.WaitGroup
	em.start(ctx, &wg)

	em.add(2, 1, "horse")
	em.add(4, 2, "sheep")
	em.add(1, 3, "pig")

	wg.Wait()

}
