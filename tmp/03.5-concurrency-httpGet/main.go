package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type siteInfo struct {
	site     string
	respTime time.Duration
}

func takeTime(sites []string, ch chan siteInfo) {
	var wg sync.WaitGroup

	for _, v := range sites {
		wg.Add(1)
		go func() {
			t := time.Now()
			_, err := http.Get(v)
			if err != nil {
				panic(err)
			}
			si := siteInfo{
				site:     v,
				respTime: time.Since(t),
			}
			ch <- si
			wg.Done()
		}()

		wg.Wait()
	}
	close(ch)
}

func main() {
	sites := []string{"https://vg.no", "https://fa.no", "https://nettavisen.no"}
	ch := make(chan siteInfo)

	go takeTime(sites, ch)

	for v := range ch {
		fmt.Printf("site=%v, load time=%v\n", v.site, v.respTime)
	}
}
