package dal

import (
	"errors"
	"fmt"
	"slices"
)

type Todo struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{Id: 1, Task: "Learn Go", Completed: false},
}

func ReadTodo() ([]Todo, error) {
	return todos, nil
}

func ReadTodoById(id int) (Todo, error) {
	for _, todo := range todos {
		if todo.Id == id {
			return todo, nil
		}
	}
	return Todo{}, errors.New("Failed to retrieve todo")
}

func CreateTodo(T Todo) (Todo, error) {
	index := len(todos)
	T.Id = index + 1
	fmt.Println(T, "=======================Todo created")
	// add Todo into todos
	todos = append(todos, T)
	return T, nil
}

func UpdateTodo(T Todo) (Todo, error) {
	// check if todo with id exists
	for i, todo := range todos {
		if todo.Id == T.Id {
			todos[i].Task = T.Task
			todos[i].Completed = T.Completed
			return todos[i], nil
		}
	}

	return T, errors.New("Todo with ID not found") // Fix: Change error message
}

func DeleteTodo(id int) error {
	for index, todo := range todos {
		if todo.Id == id {
			todos = slices.Delete(todos, index, index+1)
			return nil
		}
	}
	return errors.New("Todo with ID is not deleted")
}
