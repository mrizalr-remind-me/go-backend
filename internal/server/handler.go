package server

import (
	"github.com/mrizalr-remind-me/go-backend/internal/todo/api"
	"github.com/mrizalr-remind-me/go-backend/internal/todo/repository/mysql"
	"github.com/mrizalr-remind-me/go-backend/internal/todo/usecase"
)

func (s *server) SetupHandler() {
	setupTodo(s)
}

func setupTodo(s *server) {
	todoRepository := mysql.New(s.DB)
	todoUsecase := usecase.New(todoRepository)
	todoHandler := api.New(todoUsecase)

	v1 := s.App.Group("v1")
	todoRoute := v1.Group("/todos")
	todoHandler.SetupRoutes(todoRoute)
}
