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

//System contains both the 2 other main types
type System struct {
	Puller
	Pusher
}

//CopyContent will copy the data from db to file
func CopyContent(pllr Puller, pshr Pusher) error {

	myData, err := pullAll(pllr)
	if err != nil {
		log.Println("Error: PullAll: ", err)
	}
	fmt.Println("The content of myData = ", string(myData))

	// ----------------------------------------------------------

	n, err := pushAll(pshr, myData)
	if err != nil {
		log.Println("Error: pushAll: ", n, err)
	}

	return nil
}

// =============================================================

func main() {
	pllr := &Database{
		Name: "my first database",
	}
	pshr := &FileSystem{
		Name: "my first filesystem",
	}

	err := pllr.Open()
	if err != nil {
		log.Println("Error: ", err)
	}
	defer pllr.Close()

	err = pshr.Open()
	if err != nil {
		log.Println("Error: ", err)
	}
	defer pshr.Close()

	CopyContent(pllr, pshr)

}
