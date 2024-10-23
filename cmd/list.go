package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"go-crud/repo"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all todos",
	Run: func(cmd *cobra.Command, args []string) {
		error := cobra.RangeArgs(0, 0)(cmd, args) //immediatly executes the return func
		if error != nil {
			log.Fatal("unexpected number of arguments\n", error)
		}
		db, err := sql.Open("sqlite3", "./todo.db")
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
		todos, err := todoRepo.RetrieveAllTodos()
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID\tTASK\tCOMPLETED\tCREATED_AT\n")
		for i := range len(todos) {
			todo := todos[i]
			fmt.Printf("%d\t%s\t%t\t%s\n", todo.Id, todo.Description, todo.Completed, todo.CreatedAt)
		}
		// fmt.Printf("id inserted: %d\n", id)
	}, //Run
}
