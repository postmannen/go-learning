package main

import (
	"fmt"
	"sync"
)

type database struct {
	numberOfUsers int
}

func addUser(d *database) {
	mu.Lock()
	d.numberOfUsers++
	mu.Unlock()
	wg.Done()
}

var mu sync.Mutex
var wg sync.WaitGroup

func main() {

	db := &database{numberOfUsers: 0}

	count := 1000
	workers := 5
	startRange := count / workers

	go func() {
		for w := 1; w <= workers; w++ {
			wg.Add(count)
			go func(w int) {
				for i := (count * w) - startRange; i < startRange*w; i++ {
					go addUser(db)
					fmt.Println(i)
				}

			}(w)
		}
	}()

	wg.Wait()

	fmt.Println(db.numberOfUsers)

}
