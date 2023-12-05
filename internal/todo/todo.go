package todo

import (
	"github.com/google/uuid"
	"github.com/mrizalr-remind-me/go-backend/internal/model"
)

type Repository interface {
	FindByID(id uuid.UUID) (model.Todo, error)
	CreateTodo(todo model.Todo) error
}
type Usecase interface {
	CreateTodo(todo model.Todo) (model.Todo, error)
}
