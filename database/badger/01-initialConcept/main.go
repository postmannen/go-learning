package main

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

func main() {
	// A default badger config is provided in the DefaultOptions.
	// We use that, and just add where to store the files.
	opts := badger.DefaultOptions
	opts.Dir = "./db"
	opts.ValueDir = "./db"

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal("error: failed to open badger database: ", err, "\n")
	}
	defer db.Close()

	txn := db.NewTransaction(true)
	err = txn.Set([]byte("Bjørn Tore"), []byte("Svinningen"))
	if err != nil {
		log.Println("error: txn.Set failed: ", err)
	}

	err = txn.Commit()
	if err != nil {
		log.Println("error: txn.Commit failed: ", err)
	}

	// ---------------------GET VALUES-----------------------
	txr := db.NewTransaction(false)

	item, err := txr.Get([]byte("Bjørn Tore"))
	if err != nil {
		log.Println("error: txn.Get failed: ", err)
	}

	value, err := item.ValueCopy(nil)
	if err != nil {
		log.Println("error: item.Value failed: ", err)
	}

	fmt.Println(string(value))
}
