package db

import (
	"fmt"
	"go-crud/model"
)

const insertStm = "insert into todos (description, completed) values (?,?) returning id"

func InsertTodo(t model.Todo) (err error) {
	var id int
	err = DB.QueryRow(insertStm, t.Description, t.Completed).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Printf("id inserted: %d\n", id)
	return nil
}

const deleteStm = "delete from todos WHERE id = (?)"

func DeleteTodo(t model.Todo) error {
	_, err := DB.Exec(deleteStm, t.Id)
	if err != nil {
		return err
	}
	return nil
}

const updateStm = `update todos set description = (?) WHERE id = (?)`

func UpdateTodo(t model.Todo, newDescription string) error {
	_, err := DB.Exec(updateStm, newDescription, newDescription, t.Id)
	if err != nil {
		return err
	}
	return nil
}

const selectStm = `select id,description,completed from todos`

func RetrieveAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	rows, err := DB.Query(selectStm)
	if err != nil {
		return []model.Todo{}, err
	}
	defer rows.Close()
	for rows.Next() {
		todo := model.Todo{}
		err := rows.Scan(&todo.Id, &todo.Description, &todo.Completed)
		if err != nil {
			return []model.Todo{}, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
