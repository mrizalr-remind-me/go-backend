package todo

import (
	"github.com/google/uuid"
	"github.com/mrizalr-remind-me/go-backend/internal/model"
)

type Repository interface {
	FindByID(id uuid.UUID) (model.Todo, error)
	CreateTodo(todo model.Todo) error
	FindTodos() ([]model.Todo, error)
	UpdateTodo(id uuid.UUID, todo model.Todo) error
	DeleteTodo(id uuid.UUID) error
}
type Usecase interface {
	CreateTodo(todo model.Todo) (model.Todo, error)
	GetTodos() ([]model.Todo, error)
	GetTodo(id uuid.UUID) (model.Todo, error)
	UpdateTodo(id uuid.UUID, todo model.Todo) (model.Todo, error)
	DeleteTodo(id uuid.UUID) error
}
