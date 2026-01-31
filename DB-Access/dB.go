package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Accessing Database
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Querying the Database
	row, err := db.Query("SELECT id,users,password FROM users WHERE id=?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer row.Close()
	for row.Next() {
		var id int
		var users string
		var password string
		err = row.Scan(&id, &users, &password)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, users, password)
	}
}
