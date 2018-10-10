package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Data is the structure of the data we're working with
type Data struct {
	Line string
}

//------------------------------------------------------------------

//Xenia is a database which we need to pull data from
type Xenia struct {
	Host    string
	Timeout time.Duration
}

//Pull knows how to pull data
func (*Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF

	case 5:
		return errors.New("Failed returning data")

	default:
		d.Line = "some data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

//Pillar is a database which we will store data to
type Pillar struct {
	Host    string
	Timeout time.Duration
}

//Store stores the data in the database
func (*Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

//----------------------------------------------------

//System wraps Xenia and Pillar into a single system
type System struct {
	Xenia
	Pillar
}

//-----------------------------------------------------

//pull knows how to pull bulks of data from Xenia
func pull(x *Xenia, data []Data) (int, error) {
	for i := range data {
		if err := x.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// store knows how to store bulks of data into Pillar.
func store(p *Pillar, data []Data) (int, error) {
	for i := range data {
		if err := p.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Copy knows how to pull and store data from the System.
func Copy(sys *System, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(&sys.Xenia, data)
		if i > 0 {
			if _, err := store(&sys.Pillar, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func main() {
	sys := System{
		Xenia: Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Pillar: Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}
