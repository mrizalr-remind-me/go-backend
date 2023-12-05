package usecase

import (
	"github.com/google/uuid"
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

func (u *usecase) GetTodos() ([]model.Todo, error) {
	todos, err := u.repository.FindTodos()
	if err != nil {
		return []model.Todo{}, err
	}

	return todos, nil
}

func (u *usecase) GetTodo(id uuid.UUID) (model.Todo, error) {
	return u.repository.FindByID(id)
}

func (u *usecase) UpdateTodo(id uuid.UUID, todo model.Todo) (model.Todo, error) {
	_, err := u.GetTodo(id)
	if err != nil {
		return model.Todo{}, err
	}

	todo.Update()

	err = u.repository.UpdateTodo(id, todo)
	if err != nil {
		return model.Todo{}, err
	}

	return u.GetTodo(id)
}

func (u *usecase) DeleteTodo(id uuid.UUID) error {
	_, err := u.repository.FindByID(id)
	if err != nil {
		return err
	}

	return u.repository.DeleteTodo(id)
}
