package cmd

import (
	"fmt"
	"log"

	"go-crud/db"
	"go-crud/model"

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
		t := model.Todo{}
		id, err := db.InsertTodo(t)
		if err != nil {
			// log.Fatal("Error inserting into DB")
			panic(err)
		}
		fmt.Printf("id inserted: %d\n", id)
		// _, err = db.InsertTodoByDesc("ciao")
		// if err != nil {
		// 	// log.Fatal("Error inserting into DB")
		// 	panic(err)
		// }
		todos, err := db.RetrieveAllTodos()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v", todos)
		// fmt.Printf("id inserted: %d\n", id)
		defer func() {
			err := db.CloseDB()
			if err != nil {
				log.Println("Error closing DB connection")
			}
		}()
	}, //Run
}
