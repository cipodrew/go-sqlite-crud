package cmd

import (
	"database/sql"
	"log"
	"strconv"

	"go-crud/repo"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "complete todo with specified ID",
	Run: func(cmd *cobra.Command, args []string) {
		error := cobra.RangeArgs(1, 1)(cmd, args) //immediatly executes the return func
		if error != nil {
			log.Fatal("unexpected number of arguments\n", error)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("not an integer\n", err)
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
		err = todoRepo.CompleteTodo(id, true)
		if err != nil {
			panic(err)
		}
	}, //Run
}
