package router

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/vikbert/go-todo-app/api/controller"
	"github.com/vikbert/go-todo-app/repository"
)

func FastHttpRouter(todoRepo repository.TodoRepositoryInterface) {

	todoController := controller.NewTodoControllerFastHttp(todoRepo)

	router := fasthttprouter.New()
	router.POST("/todos", todoController.CreateTodo)
	router.GET("/todos", todoController.ListTodos)
	router.DELETE("/todo/:id", todoController.DeleteTodo)

	fasthttp.ListenAndServe("localhost:8080", router.Handler)
}
