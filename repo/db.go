package repo

import (
	// "database/sql"
	// "fmt"

	_ "github.com/mattn/go-sqlite3"
)

// remember "github.com/mattn/go-sqlite3" uses cgo

const tableName = "todos"

// func OpenDB() (*sql.DB, error) {
// 	fmt.Println("opening DB connection!")
// 	db, err := sql.Open("sqlite3", "./todo.db")
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("connections used: %d\n", db.Stats().InUse)
// 	return db, nil
// }

// func CreateDB() error {
// 	var db = sql.DB
// 	_, err := db.Exec("create table if not exists " + tableName +
// 		" (id integer not null primary key, description text, completed boolean, created_at int);")
// 	if err != nil {
// 		return fmt.Errorf("Error creating database: %s\n", tableName)
// 	}
// 	return nil
// }
