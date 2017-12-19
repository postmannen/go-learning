// DOC: https://golang.org/pkg/database/sql/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//1. Open connection

	//db, err := sql.Open("sqlite3", ":memory:")
  db, err := sql.Open("sqlite3", "./myDB.db") //return types = *DB, error
      //*DB er en peker til databasen som da blir lagt til variabelen db,
      // og error blir lagt til variabelen err
	checkErr(err)
	defer db.Close() //lukk databasen til slutt når funksjonen er ferdig

	//2. fail-fast if can't connect to DB

	checkErr(db.Ping())

	//3. create table

	_, err = db.Exec("create table USER (ID integer PRIMARY KEY, NAME string not null); delete from USER;")
    //return types = (Result, error)
    //exec utfører en QUERY uten å gi noen rows tilbake,
    //typisk ved opprettelse av tabeller
	checkErr(err)

	//4. insert data

	//4.1 Begin transaction
	tx, err := db.Begin() //return types = *Tx, error
    //begin starter transaction
    //Tx is an in-progress database transaction.
    //A transaction must end with a call to Commit or Rollback.
    //After a call to Commit or Rollback, all operations on the transaction fail with ErrTxDone.
	checkErr(err)

	//4.2 Prepare insert stmt.
	stmt, err := tx.Prepare("insert into USER(ID, NAME) values(?, ?)")
    //return types = (*Stmt, error)
    //Prepare creates a prepared statement for use within a transaction.
    //Stmt is a prepared statement. A Stmt is safe for concurrent use by multiple goroutines.
	checkErr(err)
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i, fmt.Sprint("user-", i)) //returns (result,error)
    //Exec executes a prepared statement with the given arguments and returns a Result summarizing the effect of the statement.
		checkErr(err)
	}

	//4.3 Commit transaction
	tx.Commit()
  //Commit commits the transaction.

	//5. Query data

	rows, err := db.Query("select * from USER")
	checkErr(err)
	defer rows.Close()

	//5.1 Iterate through result set
	for rows.Next() {  //return type bool
    //Next prepares the next result row for reading with the Scan method.
    //It returns true on success, or false if there is no next result row
    //or an error happened while preparing it.
    //Err should be consulted to distinguish between the two cases.
		var name string
		var id int
		err := rows.Scan(&id, &name) //func (rs *Rows) Scan(dest ...interface{}) error
    //Scan copies the columns in the current row into the values pointed at by dest.
    //The number of values in dest must be the same as the number of columns in Rows.
		checkErr(err)
		fmt.Printf("id=%d, name=%s\n", id, name)
	}

	//5.2 check error, if any, that were encountered during iteration
	err = rows.Err()
	checkErr(err)
}

func checkErr(err error, args ...string) {
	if err != nil {
		fmt.Println("Error")
		fmt.Println("%q: %s", err, args)
	}
}
