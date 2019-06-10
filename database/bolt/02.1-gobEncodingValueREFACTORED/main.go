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
	"strconv"

	"github.com/boltdb/bolt"
)

//Person describes a person
type Person struct {
	Name string
	Age  int
}

//putInDB will take a Person, encode it to gob, and put it into the db.
func putInDB(db *bolt.DB, p Person, key int) error {
	k := strconv.Itoa(key)
	//create a buffer with which is an io.Writer to satisfy gob.NewEncoder
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	var err error

	// ----------------- PUT SOMETHING INTO THE DB -------------------
	//Encode p into the bytes.buffer
	if err := enc.Encode(p); err != nil {
		return fmt.Errorf("error: failed to encode to gob: %v\n", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bu, err := tx.CreateBucketIfNotExists([]byte("persons"))
		if err != nil {
			return fmt.Errorf("error: create bucker failed: %v\n", err)
		}

		//Put the byte values of the buffer into the value with buf.Bytes.
		//NB: Using a static key for this test.
		err = bu.Put([]byte(k), buf.Bytes())
		if err != nil {
			return fmt.Errorf("error: put failed: %v\n", err)
		}

		//We came here with no errors :-)
		return nil
	})

	if err != nil {
		return fmt.Errorf("error: update failed: %v\n", err)
	}

	return nil
}

//getOutOfDB will read the value from the db, decode it from gob to []byte.
func getOutOfDB(db *bolt.DB, key int) (Person, error) {
	k := strconv.Itoa(key)
	var err error
	p := Person{}

	err = db.View(func(tx *bolt.Tx) error {
		//Open the bucket named "persons"
		bu := tx.Bucket([]byte("persons"))
		//Get the value of the key "1"
		v := bu.Get([]byte(k))
		if v == nil {
			return fmt.Errorf("No such key\n")
		}

		//Since gob.NewDecoder takes an io.Reader, we create a bytes buffer
		// who got a Read method, and then is also of type io.Reader.
		buf := &bytes.Buffer{}
		//Write the gob value that we Get earlier from the db into the buffer.
		_, err := buf.Write(v)
		if err != nil {
			return fmt.Errorf("error: write to buffer failed: %v\n ", err)
		}

		//Create a new decoder who will read from the buffer.
		dec := gob.NewDecoder(buf)

		//Decode from the buffer and into p.
		if err := dec.Decode(&p); err != nil {
			return fmt.Errorf("error: failed to decode from gob into p: %v\n", err)
		}

		//We came here with no errors :-)
		return nil
	})

	if err != nil {
		return p, fmt.Errorf("error: update failed: %v\n ", err)
	}
	return p, nil
}

func main() {
	var err error

	db, err := bolt.Open("./bolt.db", 0600, nil)
	if err != nil {
		log.Println("error: failed opening db: ", err)
	}

	//Creating a Person to put into the DB.
	p := Person{"Askeladden", 100}
	key := 1
	err = putInDB(db, p, key)
	if err != nil {
		log.Println("error: putInDB: ", err)
	}

	//Get'ing a Person out of the DB
	newP, err := getOutOfDB(db, key)
	if err != nil {
		log.Println("error: getOutOfDB failed: ", err)
	}

	fmt.Println("--- newP: ", newP)

}
