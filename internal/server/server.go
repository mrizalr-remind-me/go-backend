package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type server struct {
	App *fiber.App
	DB  *sqlx.DB
}

func New(db *sqlx.DB) *server {
	return &server{
		App: fiber.New(),
		DB:  db,
	}
}

func (s *server) Run() error {
	s.SetupHandler()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return s.App.Listen(fmt.Sprintf(":%s", port))
}
