package router

import (
	"github.com/labstack/echo/v4"
	"github.com/vikbert/go-todo-app/api/controller"
	"github.com/vikbert/go-todo-app/repository"
)

func EchoRouter(todoRepo repository.TodoRepositoryInterface) {
	todoController := controller.NewTodoControllerEcho(todoRepo)

	e := echo.New()
	e.GET("/todos", todoController.ListTodos)
	e.POST("/todos", todoController.CreateTodo)
	e.DELETE("/todo/:id", todoController.DeleteTodo)

	e.Start("localhost:8080")
}
