package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 65535; i++ {
		wg.Add(1)

		go func(port int, wg *sync.WaitGroup) {
			defer wg.Done()

			conn, err := net.DialTimeout("tcp", "10.0.0.26"+fmt.Sprintf("%v", port), time.Second*5)
			if err == nil {
				defer conn.Close()

				conn.SetReadDeadline(time.Now().Add(time.Second * 1))
				b := make([]byte, 65535)
				_, err = conn.Read(b)
				if err != nil {
					log.Printf("error: reading port %v: %v\n", port, err)
					return
				}

				fmt.Printf("Open, reading from port %v: %s\n", port, b)
			}
		}(i, &wg)
	}

	wg.Wait()
}
