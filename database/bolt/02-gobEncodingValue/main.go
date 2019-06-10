/*
	The purpose of this exercise is to store a Person type struct into a bolt-db
	as a gob encoded value ([]bytes).
	Then get that value out of the db again, decode it, and put it back into
	a new Person type struct.
*/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

//Person describes a person
type Person struct {
	Name string
	Age  int
}

func main() {
	db, err := bolt.Open("./bolt.db", 0600, nil)
	if err != nil {
		log.Println("error: failed opening db: ", err)
	}

	//create a buffer with which is an io.Writer to satisfy gob.NewEncoder
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)

	// ----------------- PUT SOMETHING INTO THE DB -------------------

	p := Person{"Askeladden", 100}

	//Encode p into the bytes.buffer
	if err := enc.Encode(p); err != nil {
		log.Println("error: failed to encode to gob: ", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bu, err := tx.CreateBucketIfNotExists([]byte("persons"))
		if err != nil {
			return fmt.Errorf("error: create bucker failed: %v\n", err)
		}

		//Put the byte values of the buffer into the value with buf.Bytes.
		err = bu.Put([]byte("1"), buf.Bytes())
		if err != nil {
			return fmt.Errorf("error: put failed: %v\n", err)
		}

		//We came here with no errors :-)
		return nil
	})

	if err != nil {
		log.Println("error: update failed: ", err)
	}

	// ----------------- GET SOMETHING OUT OF THE DB --------------------

	err = db.View(func(tx *bolt.Tx) error {
		//Open the bucket named "persons"
		bu := tx.Bucket([]byte("persons"))
		//Get the value of the key "1"
		v := bu.Get([]byte("1"))
		if v == nil {
			return fmt.Errorf("No such key\n")
		}

		//Since we will Get a type Person from the DB we will create a variable
		// to hold a Person.
		p := &Person{}
		//Since gob.NewDecoder takes an io.Reader, we create a bytes buffer
		// who got a Read method, and then is also of type io.Reader.
		buf := &bytes.Buffer{}
		//Write the gob value that we Get earlier from the db into the buffer.
		_, err := buf.Write(v)
		if err != nil {
			log.Println("error: write to buffer failed: ", err)
		}

		//Create a new decoder who will read from the buffer.
		dec := gob.NewDecoder(buf)

		//Decode from the buffer and into p.
		if err := dec.Decode(p); err != nil {
			log.Println("error: failed to decode from gob into p: ", err)
		}

		fmt.Println("--------Content of p = ", p)

		//We came here with no errors :-)
		return nil
	})

	if err != nil {
		log.Println("error: update failed: ", err)
	}

}
