package cmd

import (
	"log"
	"strconv"

	"go-crud/db"

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
		err := db.OpenDB()
		if err != nil {
			panic(err)
		}
		defer func() {
			err := db.CloseDB()
			if err != nil {
				log.Println("Error closing DB connection")
			}
		}()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("not an integer\n", err)
		}
		err = db.CompleteTodo(id, true)
		if err != nil {
			panic(err)
		}
	}, //Run
}
