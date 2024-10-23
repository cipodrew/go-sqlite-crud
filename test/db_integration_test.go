package test

import (
	"database/sql"
	"go-crud/model"
	"go-crud/repo"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInsertTodo_Success(t *testing.T) {
	db, err := sql.Open("sqlite3", "../test.db")
	if err != nil {
		t.Fatalf("error connecting to DB %#q", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			t.Fatalf("Error closing DB connection %#q", err)
		}
	}()
	repo := repo.NewTodoRepo(db)
	todo := model.Todo{
		Description: "Test task",
		Completed:   false,
	}

	id, err := repo.InsertTodo(todo)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if id <= 0 {
		t.Errorf("Expected positive ID, got %v", id)
	}
}

func TestInsertTodo_Error(t *testing.T) {
	db, err := sql.Open("sqlite3", "fail.db")
	if err != nil {
		t.Fatalf("error connecting to DB %#q", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			t.Fatalf("Error closing DB connection %#q", err)
		}
	}()
	repo := repo.NewTodoRepo(db)

	todo := model.Todo{
		Description: "Test task",
		Completed:   false,
	}
	id, err := repo.InsertTodo(todo)

	if err == nil {
		t.Errorf("Expected error, got none")
	}

	if id != -1 {
		t.Errorf("Expected ID -1, got %v", id)
	}
}
