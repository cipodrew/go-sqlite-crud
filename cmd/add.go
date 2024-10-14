package cmd

import (
	"fmt"
	"log"

	"go-crud/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a todo",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("Let's add!")
		db.OpenDB()
		err := db.CreateDB()
		if err != nil {
			panic(err)
		}
		defer func() {
			err := db.CloseDB()
			if err != nil {
				log.Println("Error closing DB connection")
			}
		}()
	}, //Run
}
