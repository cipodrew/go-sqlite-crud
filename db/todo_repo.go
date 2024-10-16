package db

import (
	"go-crud/model"
)

const insertStm = "insert into todos (description, completed) values (?,?) returning id"

func InsertTodo(t model.Todo) (insertedId int, err error) {
	var id int
	err = DB.QueryRow(insertStm, t.Description, t.Completed).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func InsertTodoByDesc(description string) (insertedId int, err error) {
	var id int
	err = DB.QueryRow(insertStm, description, false).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

const deleteStm = "delete from todos WHERE id = (?)"

func DeleteTodo(t model.Todo) error {
	_, err := DB.Exec(deleteStm, t.Id)
	if err != nil {
		return err
	}
	return nil
}
func DeleteTodoById(id int) error {
	_, err := DB.Exec(deleteStm, id)
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

const completeStm = `update todos set completed = (?) WHERE id = (?)`

func CompleteTodo(id int, completed bool) error {
	_, err := DB.Exec(completeStm, completed, id)
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
