package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// findFastestWeb will take a slice of url's represented as a []string,
// and return back the response of the webserver that returned the
// complete page first on the resp channel.
func findFastestWeb(url []string, ch chan respData) {
	var found bool
	var mu sync.Mutex // mutex to protect the first to update the found variable.

	// Loop over all the url's given as input, and start a go routine
	// with a http.Get for each of them.
	for _, u := range url {
		go func(u string) {
			t := time.Now()
			resp, err := http.Get(u)
			if err != nil {
				log.Println("error: http.Get for one", err)
			}

			if !found {
				mu.Lock()
				found = true
				mu.Unlock()

				totalTime := time.Until(t)

				ch <- respData{time: totalTime, resp: resp}
			}
		}(u)

	}
}

type respData struct {
	time time.Duration
	resp *http.Response
}

func main() {
	// respCh will contain all the data for the fastest response.
	respCh := make(chan respData, 1)

	urls := []string{"https://dagbladet.no", "https://aftenposten.no", "https://vg.no"}

	findFastestWeb(urls, respCh)

	r := <-respCh
	defer r.resp.Body.Close()

	fmt.Printf("%v : finalized first, and in time it took %v\n", r.resp.Request.URL, r.time)
}
