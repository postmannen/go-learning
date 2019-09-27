/*
The purpose of this test is to learn more on context and
cancelation of http request and responses.
In real life this example could have been done a lot easier
without waitgroup's etc, but they are put in here to be able
to see that the context done channels shut's down all the
go routines.
*/
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
	// There is actually no need to use a waitgroup in this example,
	// but it is used to wait for all goroutines to be canceled
	// by the context.
	var wg sync.WaitGroup
	fastestCh := make(chan webResponse, 1)

	for _, url := range urls {
		wg.Add(1)

		// Start a goroutine to process the the request, and read the response
		// for each url.
		go func(url string) {
			defer wg.Done()

			fmt.Printf("Starting download of %v\n", url)
			tNow := time.Now()

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				log.Println("error: NewRequestWithContext: ", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println("error: DefaultClient: ", err)
			}

			// Instead of using ioutils.Readall we have to divide up the reading
			// in smaller chunk's so we are able to check if there have been
			// sent a cancelation signal inbetween, and then exit the whole
			// go routine.
			body := []byte{}
			b := make([]byte, 64)

			var loopExit bool

			for {
				// Check if we have received a cancelation signal,
				// if not read a chunk of data from the response.Body.
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

			_ = fmt.Sprint(body)

			tSince := time.Since(tNow)

			fastestCh <- webResponse{
				server: resp.Request.URL.Hostname(),
				time:   tSince,
			}

		}(url)
	}

	f := <-fastestCh
	// We have found our fastest web server, send the cancel signal
	// to all the RequestsWithContext.
	cancel()

	fmt.Printf("\n*** Fastest was %v, and it took %v ***\n\n", f.server, f.time)

	// We want to wait to see that all the other requests also are canceled.
	wg.Wait()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	urls := []string{"https://vg.no", "https://dagbladet.no", "https://dl.google.com/go/go1.13.1.src.tar.gz", "https://digi.no", "https://itavisen.no", "https://wikipedia.no", "https://facebook.com", "https://twitter.com", "https://dn.no"}

	getFastestWeb(ctx, cancel, urls)
}
