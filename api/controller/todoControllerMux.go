package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vikbert/go-todo-app/model"
	"github.com/vikbert/go-todo-app/repository"
	"net/http"
)

type TodoControllerMux struct {
	todoRepo repository.TodoRepositoryInterface
}

func NewTodoControllerMux(repository repository.TodoRepositoryInterface) *TodoControllerMux {
	return &TodoControllerMux{
		todoRepo: repository,
	}
}

func (ctrl *TodoControllerMux) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo model.Todo
	json.NewDecoder(r.Body).Decode(&newTodo)

	entry, _ := ctrl.todoRepo.Create(newTodo)

	responseJSON(w, http.StatusCreated, entry)
}

func (ctrl *TodoControllerMux) ListTodos(w http.ResponseWriter, r *http.Request) {
	entries, _ := ctrl.todoRepo.List()

	responseJSON(w, http.StatusOK, entries)
}

func (ctrl *TodoControllerMux) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctrl.todoRepo.Delete(vars["id"])

	responseJSON(w, http.StatusAccepted, nil)
}

func responseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
