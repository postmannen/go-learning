package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func dbUpdate(db *bolt.DB, key string, value string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		//Create a bucket
		bu, err := tx.CreateBucketIfNotExists([]byte("firstBucket"))
		if err != nil {
			return fmt.Errorf("error: CreateBuckerIfNotExists failed: %v", err)
		}

		//Put a value into the bucket.
		if err := bu.Put([]byte(key), []byte(value)); err != nil {
			return err
		}

		//If all was ok, we should return a nil for a commit to happen. Any error
		// returned will do a rollback.
		return nil
	})
	return err
}

func dbView(db *bolt.DB, key string) (string, error) {
	var value string
	//View is a help function to get values out of the database.
	err := db.View(func(tx *bolt.Tx) error {
		//Open a bucket to get key's and values from.
		bu := tx.Bucket([]byte("firstBucket"))

		v := bu.Get([]byte(key))
		if len(v) == 0 {
			return fmt.Errorf("info: view: key not found")
		}

		value = string(v)

		return nil
	})

	return value, err

}

func main() {
	var err error
	//open takes the file name, permissions for that file, and database options.
	db, err := bolt.Open("./db", 0600, nil)
	if err != nil {
		log.Println("error: bolt.Open: ", err)
	}

	if err != nil {
		log.Println("error: open failed: ", err)
	}

	err = dbUpdate(db, "someKey", "someValue")
	if err != nil {
		log.Println("error: update failed: ", err)
	}

	v, err := dbView(db, "someKey2")
	if err != nil {
		log.Println("error: view failed: ", err)
	}

	fmt.Println("The value got from db = ", v)

}
