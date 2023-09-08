package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vikbert/go-todo-app/model"
	"github.com/vikbert/go-todo-app/repository"
)

type TodoControllerFiber struct {
	todoRepo repository.TodoRepositoryInterface
}

func NewTodoControllerFiber(repository repository.TodoRepositoryInterface) *TodoControllerFiber {
	return &TodoControllerFiber{
		todoRepo: repository,
	}
}

func (ctrl *TodoControllerFiber) CreateTodo(c *fiber.Ctx) error {
	newTodo := new(model.Todo)

	if err := c.BodyParser(newTodo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	entry, _ := ctrl.todoRepo.Create(*newTodo)

	return c.Status(201).JSON(entry)
}

func (ctrl *TodoControllerFiber) ListTodos(c *fiber.Ctx) error {
	entries, _ := ctrl.todoRepo.List()

	return c.Status(200).JSON(entries)
}

func (ctrl *TodoControllerFiber) DeleteTodo(c *fiber.Ctx) error {
	todoId := c.Params("id")
	ctrl.todoRepo.Delete(todoId)

	return c.Status(202).JSON(nil)
}
