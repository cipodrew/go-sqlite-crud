package cmd

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go-crud/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve an HTTP server with CRUD API",
	Run: func(cmd *cobra.Command, args []string) {
		error := cobra.RangeArgs(0, 0)(cmd, args) //immediatly executes the return func
		if error != nil {
			log.Fatal("unexpected number of arguments\n", error)
		}

		http.HandleFunc("/hi", handleHelloWorld)
		http.HandleFunc("/headers", handleGetHeaders)
		http.HandleFunc("/", handleFetchAllTodos)
		http.HandleFunc("/complete", handleCompleteTodo)
		err := http.ListenAndServe(":8090", nil)
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			log.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}, //Run
}

func handleHelloWorld(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "Hello World!")
}

func handleGetHeaders(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func handleFetchAllTodos(w http.ResponseWriter, req *http.Request) {
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
	fmt.Println(req.UserAgent()) //log user agent
	todos, err := db.RetrieveAllTodos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "ID\tTASK\tCOMPLETED\tCREATED_AT\n")
	for i := range len(todos) {
		todo := todos[i]
		fmt.Fprintf(w, "%d\t%s\t%t\t%s\n", todo.Id, todo.Description, todo.Completed, todo.CreatedAt)
	}
}

func handleCompleteTodo(w http.ResponseWriter, req *http.Request) {
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
	fmt.Println(req.UserAgent()) //log user agent
	idParam := req.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	err = db.CompleteTodo(id, true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "todo %d Completed\n", id)

}
