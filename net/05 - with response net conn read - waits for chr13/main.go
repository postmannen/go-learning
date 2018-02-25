package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//listen for inncomming connections
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error : Listener : ", err)
	}
	defer listener.Close()

	for {
		//Listen for new individual connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error : conn : ", err)
			os.Exit(1)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	blockSize := 4
	buf := make([]byte, blockSize)
	readData := []byte{}

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Error: Read: ", err)
			break
		}

		fmt.Println("Read a block of ", blockSize, " bytes , and read ", n, "bytes", ", buf contains = ", buf)

		//append each 4 byte slice to a new buffer for printing out later when iteration is done
		readData = append(readData, buf...)

		//range the buffer to check if enter was pressed in message
		foundNL := false
		for _, v := range buf {
			if v == 13 {
				foundNL = true
				break
			}
		}
		if foundNL {
			break
		}

	}
	fmt.Println("The 4 byte chunchs appended to one slice of byte contains = ", string(readData), "slice looks like = ", readData)
}
