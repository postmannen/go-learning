package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

func main() {
	//open takes the file name, permissions for that file, and database options.
	db, err := bolt.Open("./db", 0600, nil)
	if err != nil {
		log.Println("error: bolt.Open: ", err)
	}

	//Update is a helper function in the bolt package.
	// It takes care of commiting if everything went ok,
	// and it will also do a rollback if an error happened.
	err = db.Update(func(tx *bolt.Tx) error {
		//Create a bucket
		bu, err := tx.CreateBucketIfNotExists([]byte("firstBucket"))
		if err != nil {
			return fmt.Errorf("error: CreateBuckerIfNotExists failed: %v \n", err)
		}

		//Put a value into the bucket.
		if err := bu.Put([]byte("firstKey"), []byte("firstValue")); err != nil {
			return err
		}

		//If all was ok, we should return a nil for a commit to happen. Any error
		// returned will do a rollback.
		return nil
	})

	if err != nil {
		log.Println("error: db.Update failed: ", err)
	}

	//View is a help function to get values out of the database.
	err = db.View(func(tx *bolt.Tx) error {
		//Open a bucket to get key's and values from.
		bu := tx.Bucket([]byte("firstBucket"))

		v := bu.Get([]byte("firstKey"))
		fmt.Println("Value for the key=firstKey is =", string(v))

		return nil
	})

	if err != nil {
		log.Println("error: db.View failed: ", err)
	}

	// ------------------------------------------------------

	err = db.Update(func(tx *bolt.Tx) error {
		bu, err := tx.CreateBucketIfNotExists([]byte("secondBucket"))
		if err != nil {
			return fmt.Errorf("error: CreateBuckerIfNotExists failed: %v \n", err)
		}

		for i := 0; i < 20; i++ {
			err := bu.Put([]byte("key"+strconv.Itoa(i)), []byte("Value"+strconv.Itoa(i)))
			if err != nil {
				return fmt.Errorf("error: Put failed: %v\n", err)
			}
		}

		return nil
	})

	if err != nil {
		log.Println("error: db.Update with for loop failed: ", err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		bu := tx.Bucket([]byte("secondBucket"))
		cursor := bu.Cursor()

		k, v := cursor.First()
		fmt.Println("first key/value = ", k, v)

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		//for {
		//	k, v := cursor.Next()
		//	if k != nil {
		//		//If the key = nil then we have looped all the items and break out.
		//		log.Println("k = nil")
		//		break
		//	}
		//	fmt.Println("the rest of the key/values = ", string(k), string(v))
		//}

		return nil
	})

	if err != nil {
		log.Println("error: db.View with for loop failed: ", err)
	}
}
