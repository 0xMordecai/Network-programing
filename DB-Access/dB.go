package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	row, err := db.Query("SELECT id,users,password FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer row.Close()
	defer db.Close()
}
