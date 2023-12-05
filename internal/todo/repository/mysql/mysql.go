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
	if err != nil {
		return model.Todo{}, fmt.Errorf("error scanning query result - %v", err)
	}
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

func (r *repository) FindTodos() ([]model.Todo, error) {
	query := `SELECT id, title, description, remind_at, created_at, updated_at
	FROM todo`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing query - %v", err)
	}

	todos := []model.Todo{}
	rows, err := stmt.Query()
	for rows.Next() {
		todo := model.Todo{}
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.RemindAt,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning query result - %v", err)
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *repository) UpdateTodo(id uuid.UUID, todo model.Todo) error {
	query := `UPDATE todo
	SET
		title = :title,
		description = :description,
		remind_at = :remind_at,
		updated_at = :updated_at
	WHERE id = :id`

	params := map[string]any{
		"id":          id,
		"title":       todo.Title,
		"description": todo.Description,
		"remind_at":   todo.RemindAt,
		"updated_at":  todo.UpdatedAt,
	}

	nstmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return fmt.Errorf("error preparing query - %v", err)
	}

	_, err = nstmt.Exec(params)
	return err
}

func (r *repository) DeleteTodo(id uuid.UUID) error {
	query := `DELETE
	FROM todo
	WHERE id = :id`

	params := map[string]any{"id": id}

	nstmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return fmt.Errorf("error preparing query - %v", err)
	}

	_, err = nstmt.Exec(params)
	return err
}
