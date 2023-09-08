package router

import (
	"github.com/gorilla/mux"
	"github.com/vikbert/go-todo-app/api/controller"
	"github.com/vikbert/go-todo-app/repository"
	"net/http"
)

func MuxRouter(todoRepo repository.TodoRepositoryInterface) {

	todoController := controller.NewTodoControllerMux(todoRepo)

	router := mux.NewRouter()
	router.HandleFunc("/todos", todoController.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", todoController.ListTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", todoController.DeleteTodo).Methods("DELETE")

	http.ListenAndServe("localhost:8080", router)
}
