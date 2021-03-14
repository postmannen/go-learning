package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	ctx, cancel := context.WithCancel(context.Background())

	req, err := http.NewRequestWithContext(ctx, "post", "https://ac0bfe16f998574acf0c84426d38d18b2b.endpoint.twilio.com", nil)
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
}
