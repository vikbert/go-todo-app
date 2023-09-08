package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vikbert/go-todo-app/api/controller"
	"github.com/vikbert/go-todo-app/repository"
)

func FiberRouter(todoRepo repository.TodoRepositoryInterface) {
	todoController := controller.NewTodoControllerFiber(todoRepo)

	app := fiber.New()
	app.Get("/todos", todoController.ListTodos)
	app.Post("/todos", todoController.CreateTodo)
	app.Delete("/todo/:id", todoController.DeleteTodo)

	app.Listen("localhost:8080")
}
