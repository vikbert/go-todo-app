package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/vikbert/go-todo-app/api/controller"
	"github.com/vikbert/go-todo-app/repository"
	"net/http"
)

func ChiRouter(todoRepo repository.TodoRepositoryInterface) {
	todoController := controller.NewTodoControllerChi(todoRepo)

	router := chi.NewRouter()
	router.Get("/todos", todoController.ListTodos)
	router.Post("/todos", todoController.CreateTodo)
	router.Delete("/todo/{id}", todoController.DeleteTodo)

	http.ListenAndServe("localhost:8080", router)
}
