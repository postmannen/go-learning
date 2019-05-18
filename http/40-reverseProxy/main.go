//Simple reverse proxy
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//forwarders is a map where the key is the internal URL's, and the Key is the external URL.
type forwarders map[string]string

func forwarder(f forwarders) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		extUrl, ok := f[r.Host]
		if !ok {
			fmt.Fprintf(w, "no forwarder for domain found, %v\n", ok)
			return
		}

		resp, err := http.Get(extUrl + r.URL.RequestURI())
		if err != nil {
			log.Printf("error: http.Get: %v\n", err)
		}
		defer resp.Body.Close()

		//Copy all the header values into the ResponseWriter.
		for k, v := range resp.Header {
			w.Header()[k] = v
		}
		w.WriteHeader(resp.StatusCode)

		//Copy the HTTP body into the ResponseWriter.
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			log.Printf("error: io.Copy: %v\n", err)
		}

	}
}

func main() {
	forwarders := map[string]string{
		"nrk.localhost":   "https://nrk.no",
		"vg.localhost":    "https://vg.no",
		"erter.localhost": "https://erter.org",
	}

	http.HandleFunc("/", forwarder(forwarders))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Printf("error: http.ListenAndServe: %v\n", err)
	}
}
