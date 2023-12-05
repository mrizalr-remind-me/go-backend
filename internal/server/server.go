package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	App *fiber.App
}

func New() *server {
	return &server{
		App: fiber.New(),
	}
}

func (s *server) Run() error {
	s.SetupHandler()

	port := os.Getenv("SERVER_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return s.App.Listen(fmt.Sprintf(":%s", port))
}
