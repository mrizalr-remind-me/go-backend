package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *server) SetupHandler() {
	s.App.Get("health", func(c *fiber.Ctx) error {
		return c.
			Status(http.StatusOK).
			JSON(fiber.Map{"status": "ok"})
	})
}
