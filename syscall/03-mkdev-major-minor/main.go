package main

import (
	"log"
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
	devPath := filepath.Join("/tmp/ttys191")

	{
		_, err := os.Stat(devPath)
		if err != nil {
			log.Printf("error: os.Stat: %v\n", err)
		}
	}

	dev := unix.Mkdev(4, 255)

	err := syscall.Mknod(devPath, syscall.S_IFCHR|0666, int(dev))
	if err != nil {
		log.Printf("error: mknod: %v\n", err)
	}
}
