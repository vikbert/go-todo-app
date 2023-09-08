package repository

import "github.com/vikbert/go-todo-app/model"

type TodoRepositoryInterface interface {
	Create(todo model.Todo) (model.Todo, error)
	List() ([]model.Todo, error)
	Delete(id string)
}
