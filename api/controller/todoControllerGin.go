package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vikbert/go-todo-app/model"
	"github.com/vikbert/go-todo-app/repository"
	"net/http"
)

type TodoControllerGin struct {
	todoRepo repository.TodoRepositoryInterface
}

func NewTodoControllerGin(repository repository.TodoRepositoryInterface) *TodoControllerGin {
	return &TodoControllerGin{
		todoRepo: repository,
	}
}

func (ctrl *TodoControllerGin) CreateTodo(c *gin.Context) {
	var newTodo model.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	entry, _ := ctrl.todoRepo.Create(newTodo)

	c.JSON(http.StatusCreated, entry)
}

func (ctrl *TodoControllerGin) ListTodos(c *gin.Context) {

	entries, _ := ctrl.todoRepo.List()

	c.JSON(http.StatusOK, entries)
}

func (ctrl *TodoControllerGin) DeleteTodo(c *gin.Context) {
	todoId := c.Param("id")
	ctrl.todoRepo.Delete(todoId)

	c.JSON(http.StatusAccepted, nil)
}
