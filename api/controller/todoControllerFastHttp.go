package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12/httptest"
	"github.com/valyala/fasthttp"
	"github.com/vikbert/go-todo-app/model"
	"github.com/vikbert/go-todo-app/repository"
)

type TodoControllerFastHttp struct {
	todoRepo repository.TodoRepositoryInterface
}

func NewTodoControllerFastHttp(repository repository.TodoRepositoryInterface) *TodoControllerFastHttp {
	return &TodoControllerFastHttp{
		todoRepo: repository,
	}
}

func (ctrl *TodoControllerFastHttp) CreateTodo(c *fasthttp.RequestCtx) {
	var newTodo model.Todo
	err := json.Unmarshal(c.PostBody(), &newTodo)
	if err != nil {
		c.Error("Failed to parse request body", 500)
		return
	}

	entry, _ := ctrl.todoRepo.Create(newTodo)

	body, err := json.Marshal(entry)
	if err != nil {
		c.Error("Failed to parse the response body", 500)
		return
	}

	c.SetContentType("application/json")
	c.SetStatusCode(201)
	c.SetBody(body)
}

func (ctrl *TodoControllerFastHttp) ListTodos(c *fasthttp.RequestCtx) {
	entries, _ := ctrl.todoRepo.List()

	body, err := json.Marshal(entries)
	if err != nil {
		c.Error("Failed to parse response data", 500)
		return
	}

	c.SetContentType("application/json")
	c.SetBody(body)
}

type TodoRequest struct {
	Id string
}

func (ctrl *TodoControllerFastHttp) DeleteTodo(c *fasthttp.RequestCtx) {
	todoId := c.UserValue("id")
	ctrl.todoRepo.Delete(fmt.Sprintf("%v", todoId))

	c.SetStatusCode(httptest.StatusAccepted)
}
