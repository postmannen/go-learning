package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type webResponse struct {
	server string
	time   time.Duration
}

func getFastestWeb(ctx context.Context, cancel context.CancelFunc, urls []string) {
	var wg sync.WaitGroup
	fastestCh := make(chan webResponse, 1)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			fmt.Printf("Starting download of %v\n", url)
			tNow := time.Now()

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				log.Println(err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println(err)
			}

			body := []byte{}
			b := make([]byte, 64)

			var loopExit bool

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("received done for url = %v, and returning out early, no further downloading of page\n", url)
					return
				default:
					n, err := resp.Body.Read(b)
					if err == io.EOF && n == 0 {
						loopExit = true
						break
					}

					body = append(body, b...)
				}
				if loopExit {
					break
				}
			}

			//body, err := ioutil.ReadAll(resp.Body)
			//if err != nil {
			//	log.Println(err)
			//}

			_ = fmt.Sprint(body)

			tSince := time.Since(tNow)

			fastestCh <- webResponse{
				server: resp.Request.URL.Hostname(),
				time:   tSince,
			}
			log.Printf("Before calling cancel in the goroutine for %v\n", url)
			cancel()
			log.Printf("done with go routine for %v\n", url)
		}(url)
	}

	f := <-fastestCh
	fmt.Printf("*** Fastest was %v, and it took %v ***\n\n", f.server, f.time)
	wg.Wait()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	urls := []string{"https://vg.no", "https://dagbladet.no", "https://dl.google.com/go/go1.13.1.src.tar.gz"}

	getFastestWeb(ctx, cancel, urls)
}
