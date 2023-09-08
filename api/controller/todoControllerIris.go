package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/vikbert/go-todo-app/model"
	"github.com/vikbert/go-todo-app/repository"
)

type TodoControllerIris struct {
	todoRepo repository.TodoRepositoryInterface
}

func NewTodoControllerIris(repository repository.TodoRepositoryInterface) *TodoControllerIris {
	return &TodoControllerIris{
		todoRepo: repository,
	}
}

func (ctrl *TodoControllerIris) CreateTodo(c iris.Context) {
	var newTodo = model.Todo{}

	if err := c.ReadJSON(&newTodo); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{"error": err.Error()})
		return
	}

	entry, _ := ctrl.todoRepo.Create(newTodo)

	c.StatusCode(iris.StatusCreated)
	c.JSON(entry)
}

func (ctrl *TodoControllerIris) ListTodos(c iris.Context) {
	entries, _ := ctrl.todoRepo.List()

	c.JSON(entries)
}

func (ctrl *TodoControllerIris) DeleteTodo(c iris.Context) {
	todoId := c.Params().GetString("id")
	ctrl.todoRepo.Delete(todoId)

	c.StatusCode(iris.StatusAccepted)
}
