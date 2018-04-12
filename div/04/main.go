package main

import (
	"fmt"
	"io"
	"strings"
)

//hardDisk is our disk type
type hardDisk struct {
	diskData io.Reader
}

func (h hardDisk) Read(p []byte) (n int, err error) {
	fmt.Println("using hardFisk's Read method")
	n, err = h.diskData.Read(p)
	return n, err
}

func diskReader(p io.Reader) []byte {
	tmp := []byte{}
	for {
		buf := make([]byte, 4)
		_, err := p.Read(buf)
		if err != nil {
			fmt.Println("diskReader: failed read", err)
			break
		}
		tmp = append(tmp, buf...)
	}
	return tmp

}

func main() {
	disk1 := hardDisk{diskData: strings.NewReader("disk1 apekatt")}

	tmpData := diskReader(disk1)
	fmt.Printf("disk1's type = %T \n", disk1)
	fmt.Printf("The type is %T, and the data = %v\n", tmpData, string(tmpData))
}
