package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// remember "github.com/mattn/go-sqlite3" uses cgo
var DB *sql.DB

const tableName = "todos"

func OpenDB() error {
	fmt.Println("opening connection!")
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return err
	}
	DB = db
	fmt.Printf("connections used: %d\n", db.Stats().InUse)
	return nil
}

func CloseDB() error {
	return DB.Close()
}

func CreateDB() error {
	_, err := DB.Exec("create table if not exists " + tableName +
		" (id integer not null primary key, description text, completed boolean);")
	if err != nil {
		return fmt.Errorf("Error creating database: %s\n", tableName)
	}
	return nil
}
