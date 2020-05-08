package main

import (
	"io"
	"net"

	"log"
)

func main() {

	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Printf("error: net.listen failed %v\n", err)
	}

	for {
		// Wait for new connections
		conn, err := l.Accept()
		if err != nil {
			log.Printf("error: accept failed %v\n", err)
		}
		println("Got a new connection")

		go func(c net.Conn) {
			dst, err := net.Dial("tcp", "www.vg.no:80")
			if err != nil {
				log.Println("error: dialing www.vg.no: ", err)
			}
			defer dst.Close()

			// copy the content from the client to www.vg.no
			go func() {
				_, err = io.Copy(dst, c)
				if err != nil {
					log.Println("error: io.Copy to www.vg.no: ", err)
				}
			}()

			// copy the content received from www.vg.no to the client
			_, err = io.Copy(c, dst)
			if err != nil {
				log.Println("error: io.Copy from www.vg.no to client: ", err)
			}
		}(conn)
	}

}
