package api

import "github.com/gofiber/fiber/v2"

func (h *handler) SetupRoutes(r fiber.Router) {
	r.Get("", h.GetTodos)
	r.Post("", h.CreateTodo)
}
