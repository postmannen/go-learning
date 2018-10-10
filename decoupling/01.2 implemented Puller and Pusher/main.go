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

//Puller is any type that have Pull method
type Puller interface {
	Pull([]byte) (int, error)
}

//Func will try to pull all data,
//or return an error if it fails.
func pullAll(db Puller) (d []byte, err error) {
	for {
		e := make([]byte, 1)
		n, err := db.Pull(e)
		if n == 0 && err != nil {
			return d, err
		}
		d = append(d, e...)
	}
}

// =============================================================

//FileSystem is a file system
type FileSystem struct {
	Name      string
	FSHandler *os.File
}

//Open a file system
func (fs *FileSystem) Open() (err error) {
	fs.FSHandler, err = os.Create("db2.db")
	if err != nil {
		log.Println("Failed to open file system file : ", err)
		return err
	}
	return nil
}

//Close will close the database for reading and writing
func (fs *FileSystem) Close() error {
	err := fs.FSHandler.Close()
	return err
}

//Push will push a small chunc of data to the file system
func (fs *FileSystem) Push(b []byte) (n int, err error) {
	n, err = fs.FSHandler.Write(b)
	if n == 0 && err != nil {
		fmt.Println("Error: Push: ", err)
		return n, err
	}
	return n, nil
}

//Pusher is an interface for type's with a push method
type Pusher interface {
	Push([]byte) (int, error)
}

//pushAll will push a larger chunc of data to the file system
//Since *FileSystem is a Pusher we can replace FileSystem with Pusher as input
func pushAll(fs Pusher, b []byte) (n int, err error) {
	n, err = fs.Push(b)
	if n == 0 && err != nil {
		return n, err
	}
	return n, nil
}

// =============================================================

//System is embedding both the Puller and the Pusher
type System struct {
	Puller
	Pusher
}

// =============================================================

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

	// ----------------------------------------------------------

	fs := FileSystem{Name: "my first filesystem"}
	err = fs.Open()
	if err != nil {
		log.Println("Error: ", err)
	}
	defer fs.Close()

	n, err := pushAll(&fs, myData)
	if err != nil {
		log.Println("Error: pushAll: ", n, err)
	}
}
