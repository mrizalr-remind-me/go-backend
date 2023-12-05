package model

import (
	"time"
)

type Todo struct {
	base
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	RemindAt    *time.Time `json:"remind_at" db:"remind_at"`
}
