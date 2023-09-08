package repository

import (
	"github.com/vikbert/go-todo-app/model"
	"gorm.io/gorm"
)

type TodoOrmRepository struct {
	db *gorm.DB
}

func NewTodoOrmRepository(db *gorm.DB) *TodoOrmRepository {
	return &TodoOrmRepository{db: db}
}

func (r *TodoOrmRepository) Create(todo model.Todo) (model.Todo, error) {
	result := r.db.Create(&todo)

	return todo, result.Error
}

func (r *TodoOrmRepository) List() ([]model.Todo, error) {
	var todos []model.Todo
	result := r.db.Find(&todos)

	return todos, result.Error
}

func (r *TodoOrmRepository) Delete(id string) {
	var todos []model.Todo
	r.db.Delete(&todos, id)
}
