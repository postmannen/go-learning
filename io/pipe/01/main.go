package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Pipes can be used to convert a type that is a writer into a reader.
// This can be handy when you want to use io.Copy which wants to copy
// the content of an io.Reader over into an io.Writer.
// Below we are using Fprint to write some text to a Writer. Since Fprint
// writes into an writer we can not directly use that with an io.Copy since
// io.Copy wants to copy from a reader into a writer. We can then use the
// Pipe to convert from writer to reader as below.
//
// We could have written directly to os.Stdout, but it was not done so
// to try to make an example that was as small as possible to show how
// it can be used.
func main() {
	r, w := io.Pipe()

	go func() {
		for {
			_, err := io.Copy(os.Stdout, r)
			if err != nil {
				log.Printf("error: io.Copy failed: %v\n", err)
			}
		}
	}()

	fmt.Fprintf(w, "some text\n")

	time.Sleep(time.Second * 1)

}
