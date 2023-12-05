package usecase

import (
	"github.com/mrizalr-remind-me/go-backend/internal/model"
	"github.com/mrizalr-remind-me/go-backend/internal/todo"
)

type usecase struct {
	repository todo.Repository
}

func New(repository todo.Repository) *usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) CreateTodo(todo model.Todo) (model.Todo, error) {
	todo.Create()

	err := u.repository.CreateTodo(todo)
	if err != nil {
		return model.Todo{}, err
	}

	createdTodo, err := u.repository.FindByID(todo.ID)
	if err != nil {
		return model.Todo{}, err
	}

	return createdTodo, nil

}
