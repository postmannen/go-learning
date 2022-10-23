package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	minVersion := flag.Int("minVersion", 35000, "min swupd version")
	maxVersion := flag.Int("maxVersion", 37000, "max swupd version")
	flag.Parse()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	for minV := *minVersion; minV <= *maxVersion; minV += 10 {

		version := strconv.Itoa(minV)

		ctx, cancel := context.WithCancel(context.Background())

		req, err := http.NewRequestWithContext(ctx, "get", "https://cdn.download.clearlinux.org/releases/"+version+"/clear/RELEASENOTES", nil)
		if err != nil {
			log.Printf("error: newRequest failed: %v\n", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("error: client.Do failed: %v\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			cancel()
			fmt.Printf("not 200, where %#v, bailing out\n", resp.StatusCode)
			return
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("error: ReadAll failed: %v\n", err)
		}

		fmt.Printf("%s\n", b)
		fmt.Printf("-----------------------------------------------------------------------------------\n")

	}
}
