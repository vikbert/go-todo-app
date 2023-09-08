package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vikbert/go-todo-app/model"
	"github.com/vikbert/go-todo-app/repository"
	"net/http"
)

type TodoControllerChi struct {
	todoRepo repository.TodoRepositoryInterface
}

func NewTodoControllerChi(repository repository.TodoRepositoryInterface) *TodoControllerChi {
	return &TodoControllerChi{
		todoRepo: repository,
	}
}

func (ctrl *TodoControllerChi) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo model.Todo
	json.NewDecoder(r.Body).Decode(&newTodo)

	entry, _ := ctrl.todoRepo.Create(newTodo)
	responseJson(w, http.StatusCreated, entry)
}

func (ctrl *TodoControllerChi) ListTodos(w http.ResponseWriter, r *http.Request) {
	entries, _ := ctrl.todoRepo.List()

	responseJson(w, http.StatusOK, entries)
}

func (ctrl *TodoControllerChi) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "id")

	ctrl.todoRepo.Delete(todoId)

	responseJson(w, http.StatusAccepted, nil)
}

func responseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
