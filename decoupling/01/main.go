package main

import (
	"fmt"
	"log"
	"os"
)

//Database is a Database
type Database struct {
	Name      string
	DBHandler *os.File
}

//Open a database
func (db *Database) Open() (err error) {
	db.DBHandler, err = os.Open("db1.db")
	if err != nil {
		log.Println("Failed to open DB file : ", err)
		return err
	}
	return nil
}

//Close will close the database for reading and writing
func (db *Database) Close() error {
	err := db.DBHandler.Close()
	return err
}

//Pull will pull a chunc of data from the database
func (db *Database) Pull(b []byte) (int, error) {
	n, err := db.DBHandler.Read(b)
	return n, err
}

//Func will try to pull all data,
//or return an error if it fails.
func pullAll(db *Database) (d []byte, err error) {
	for {
		e := make([]byte, 1)
		n, err := db.Pull(e)
		if n == 0 && err != nil {
			return d, err
		}
		d = append(d, e...)
		fmt.Println("The content of d = ", d)
		fmt.Println("The content of e = ", d)
	}
}

func main() {
	db := &Database{Name: "my first database"}
	err := db.Open()
	if err != nil {
		log.Println("Error: ", err)
	}
	defer db.Close()

	myData, err := pullAll(db)
	if err != nil {
		log.Println("Error: PullAll: ", err)
	}
	fmt.Println("The content of myData = ", string(myData))
}
