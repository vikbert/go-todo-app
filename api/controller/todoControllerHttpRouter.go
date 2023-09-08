package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/vikbert/go-todo-app/model"
	"github.com/vikbert/go-todo-app/repository"
	"net/http"
)

type TodoControllerHttpRouter struct {
	todoRepo repository.TodoRepositoryInterface
}

func NewTodoControllerHttpRouter(repository repository.TodoRepositoryInterface) *TodoControllerHttpRouter {
	return &TodoControllerHttpRouter{
		todoRepo: repository,
	}
}

func (ctrl *TodoControllerHttpRouter) CreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newTodo model.Todo
	json.NewDecoder(r.Body).Decode(&newTodo)

	entry, _ := ctrl.todoRepo.Create(newTodo)

	returnJSON(w, http.StatusCreated, entry)
}

func (ctrl *TodoControllerHttpRouter) ListTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	entries, _ := ctrl.todoRepo.List()

	returnJSON(w, http.StatusOK, entries)
}

func (ctrl *TodoControllerHttpRouter) DeleteTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	todoId := ps.ByName("id")

	ctrl.todoRepo.Delete(todoId)

	returnJSON(w, http.StatusAccepted, nil)
}

func returnJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
