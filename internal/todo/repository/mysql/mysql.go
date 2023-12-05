package mysql

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mrizalr-remind-me/go-backend/internal/model"
)

type repository struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *repository {
	return &repository{
		DB: DB,
	}
}

func (r *repository) FindByID(id uuid.UUID) (model.Todo, error) {
	query := `SELECT 
	id, title, description, remind_at, created_at, updated_at
	FROM todo
	WHERE id = :id`

	params := map[string]any{"id": id}

	nstmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return model.Todo{}, fmt.Errorf("error preparing query - %v", err)
	}

	todo := new(model.Todo)
	err = nstmt.QueryRow(params).StructScan(todo)
	return *todo, nil
}

func (r *repository) CreateTodo(todo model.Todo) error {
	query := `INSERT INTO todo(id, title, description, remind_at, created_at)
	VALUES(:id, :title, :description, :remind_at, :created_at)`

	params := map[string]any{
		"id":          todo.ID,
		"title":       todo.Title,
		"description": todo.Description,
		"remind_at":   todo.RemindAt,
		"created_at":  todo.CreatedAt,
	}

	nstmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return fmt.Errorf("error preparing query - %v", err)
	}

	_, err = nstmt.Exec(params)
	return err
}
