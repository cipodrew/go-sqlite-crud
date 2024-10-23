package cmd

import (
	"fmt"
	"log"

	"go-crud/repo"

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

		db, err := repo.OpenDB()
		if err != nil {
			log.Fatal("error connecting to DB")
		}
		defer func() {
			err := db.Close()
			if err != nil {
				log.Println("Error closing DB connection")
			}
		}()

		todoRepo := repo.NewTodoRepo(db)
		fmt.Printf("Let's add!\n")
		// t := model.Todo{}
		// id, err := db.InsertTodo(t)
		// if err != nil {
		// 	// log.Fatal("Error inserting into DB")
		// 	panic(err)
		// }
		id, err := todoRepo.InsertTodoByDesc(args[0])
		if err != nil {
			// log.Fatal("Error inserting into DB")
			panic(err)
		}
		fmt.Printf("id inserted: %d\n", id)
		todos, err := todoRepo.RetrieveAllTodos()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v", todos)
		// fmt.Printf("id inserted: %d\n", id)
	}, //Run
}
