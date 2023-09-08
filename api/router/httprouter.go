package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/vikbert/go-todo-app/api/controller"
	"github.com/vikbert/go-todo-app/repository"
	"net/http"
)

func HttpRouterRouter(todoRepo repository.TodoRepositoryInterface) {

	todoController := controller.NewTodoControllerHttpRouter(todoRepo)

	router := httprouter.New()
	router.POST("/todos", todoController.CreateTodo)
	router.GET("/todos", todoController.ListTodos)
	router.DELETE("/todo/:id", todoController.DeleteTodo)

	http.ListenAndServe("localhost:8080", router)
}
