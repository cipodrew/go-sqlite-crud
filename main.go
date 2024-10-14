package main

import (
	"fmt"
	"go-crud/cmd"
	"log"
)

func main() {
	fmt.Println("starting")
	openDB()
	err := createDB()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := closeDB()
		if err != nil {
			log.Println("Error closing DB connection")
		}
	}()

	cmd.Execute()
}
