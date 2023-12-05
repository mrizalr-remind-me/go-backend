package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr-remind-me/go-backend/internal/model"
	"github.com/mrizalr-remind-me/go-backend/internal/todo"
)

type handler struct {
	usecase todo.Usecase
}

func New(usecase todo.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) CreateTodo(c *fiber.Ctx) error {
	payload := new(model.Todo)
	err := c.BodyParser(payload)
	if err != nil {
		return err
	}

	createdTodo, err := h.usecase.CreateTodo(*payload)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(createdTodo)
}

func (h *handler) GetTodos(c *fiber.Ctx) error {
	todos, err := h.usecase.GetTodos()
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(todos)
}
