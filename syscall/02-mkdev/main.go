package main

import (
	"log"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
	dev := unix.Mkdev(4, 255)

	err := syscall.Mknod("/tmp/ttys191", syscall.S_IFCHR|0666, int(dev))
	if err != nil {
		log.Printf("error: mknod: %v\n", err)
	}
}
