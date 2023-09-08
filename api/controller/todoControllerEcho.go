package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/vikbert/go-todo-app/model"
	"github.com/vikbert/go-todo-app/repository"
	"net/http"
)

type TodoControllerEcho struct {
	todoRepo repository.TodoRepositoryInterface
}

func NewTodoControllerEcho(repository repository.TodoRepositoryInterface) *TodoControllerEcho {
	return &TodoControllerEcho{
		todoRepo: repository,
	}
}

func (ctrl *TodoControllerEcho) CreateTodo(c echo.Context) error {
	newTodo := new(model.Todo)

	if err := c.Bind(newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	entry, _ := ctrl.todoRepo.Create(*newTodo)
	return c.JSON(http.StatusCreated, entry)
}

func (ctrl *TodoControllerEcho) ListTodos(c echo.Context) error {
	entries, _ := ctrl.todoRepo.List()

	return c.JSON(http.StatusOK, entries)
}

func (ctrl *TodoControllerEcho) DeleteTodo(c echo.Context) error {
	todoId := c.Param("id")
	ctrl.todoRepo.Delete(todoId)

	return c.JSON(http.StatusAccepted, nil)
}
