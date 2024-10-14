package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenDB() error {
	fmt.Println("opening connection!")
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return err
	}
	fmt.Printf("connections used: %d", db.Stats().InUse)
	return nil
}
