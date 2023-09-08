package router

import (
	"github.com/kataras/iris/v12"
	"github.com/vikbert/go-todo-app/api/controller"
	"github.com/vikbert/go-todo-app/repository"
)

func IrisRouter(todoRepo repository.TodoRepositoryInterface) {
	todoController := controller.NewTodoControllerIris(todoRepo)

	app := iris.New()
	app.Get("/todos", todoController.ListTodos)
	app.Post("/todos", todoController.CreateTodo)
	app.Delete("/todo/{id:string}", todoController.DeleteTodo)

	app.Listen("localhost:8080")
}
