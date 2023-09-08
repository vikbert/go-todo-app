package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vikbert/go-todo-app/api/controller"
	"github.com/vikbert/go-todo-app/repository"
)

func GinRouter(todoRepo repository.TodoRepositoryInterface) {
	todoController := controller.NewTodoControllerGin(todoRepo)

	engine := gin.Default()
	engine.GET("/todos", todoController.ListTodos)
	engine.POST("/todos", todoController.CreateTodo)
	engine.DELETE("/todo/:id", todoController.DeleteTodo)

	engine.Run("localhost:8080")
}
