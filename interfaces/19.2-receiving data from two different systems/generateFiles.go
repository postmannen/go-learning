package main

import (
	"encoding/xml"
	"math/rand"
	"os"

	"github.com/google/uuid"
)

func createFile() error {
	const stringIDLength int = 32

	{
		f, err := os.Create("testData1.xml")
		if err != nil {
			return err
		}
		defer f.Close()

		enc := xml.NewEncoder(f)

		//int as ID
		for i := 0; i < 10; i++ {
			var v obj1
			v.ID = i
			v.Data = []byte("some data")

			enc.Indent("   ", "      ")
			err = enc.Encode(v)
			if err != nil {
				return err
			}
		}
	}

	{
		f, err := os.Create("testData2.xml")
		if err != nil {
			return err
		}
		defer f.Close()

		enc := xml.NewEncoder(f)
		//String as ID
		for i := 0; i < 10; i++ {
			var v obj2
			var b []byte
			for i := 0; i <= stringIDLength; i++ {
				rnd := rand.Intn(25) + 97
				b = append(b, byte(rnd))
			}
			v.ID = string(b)
			v.Data = []byte("some data")

			enc.Indent("   ", "      ")
			err = enc.Encode(v)
			if err != nil {
				return err
			}
		}
	}

	{
		f, err := os.Create("testData3.xml")
		if err != nil {
			return err
		}
		defer f.Close()

		enc := xml.NewEncoder(f)
		//uuid as ID
		for i := 0; i < 10; i++ {
			var v obj3
			v.ID = uuid.New()
			v.Data = []byte("some data")

			enc.Indent("   ", "      ")
			err = enc.Encode(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
