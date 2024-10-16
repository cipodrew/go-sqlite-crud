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
		error := cobra.RangeArgs(1, 1)(cmd, args) //immediatly executes the return func
		if error != nil {
			log.Fatal("unexpected number of arguments\n", error)
		}

		fmt.Printf("Let's add!\n")
		err := db.OpenDB()
		if err != nil {
			panic(err)
		}
		err = db.CreateDB()
		if err != nil {
			panic(err)
		}
		defer func() {
			err := db.CloseDB()
			if err != nil {
				log.Println("Error closing DB connection")
			}
		}()
		// t := model.Todo{}
		// id, err := db.InsertTodo(t)
		// if err != nil {
		// 	// log.Fatal("Error inserting into DB")
		// 	panic(err)
		// }
		id, err := db.InsertTodoByDesc(args[0])
		if err != nil {
			// log.Fatal("Error inserting into DB")
			panic(err)
		}
		fmt.Printf("id inserted: %d\n", id)
		todos, err := db.RetrieveAllTodos()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v", todos)
		// fmt.Printf("id inserted: %d\n", id)
	}, //Run
}
