package main

import (
	"fmt"
	"go-crud/cmd"
)

func main() {
	fmt.Println("starting")
	OpenDB()

	cmd.Execute()
}
