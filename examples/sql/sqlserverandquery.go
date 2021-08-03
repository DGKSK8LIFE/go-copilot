// Write a program that opens a mysql connection
// and fetches rows from a table using prepared statements.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Execute a prepared statement.
func execute(stmt *sql.Stmt) {
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
}

func main() {
	db, err := sql.Open("mysql", os.Getenv("SQL_DATABASE_USER")+":"+os.Getenv("SQL_DATABASE_PASSWORD")+"@tcp("+os.Getenv("SQL_HOST")+")/"+os.Getenv("SQL_DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	execute(stmt)
}
