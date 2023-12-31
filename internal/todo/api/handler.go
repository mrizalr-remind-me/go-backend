package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	createdTodo, err := h.usecase.CreateTodo(*payload)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.Status(http.StatusCreated).JSON(createdTodo)
}

func (h *handler) GetTodos(c *fiber.Ctx) error {
	todos, err := h.usecase.GetTodos()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.Status(http.StatusOK).JSON(todos)
}

func (h *handler) GetTodo(c *fiber.Ctx) error {
	todoUUID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	todo, err := h.usecase.GetTodo(todoUUID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.Status(http.StatusOK).JSON(todo)
}

func (h *handler) UpdateTodo(c *fiber.Ctx) error {
	todoUUID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	payload := new(model.Todo)
	err = c.BodyParser(payload)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	todo, err := h.usecase.UpdateTodo(todoUUID, *payload)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.Status(http.StatusOK).JSON(todo)
}

func (h *handler) DeleteTodo(c *fiber.Ctx) error {
	todoUUID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	err = h.usecase.DeleteTodo(todoUUID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("todo with id %s successfully deleted", todoUUID.String())})
}
