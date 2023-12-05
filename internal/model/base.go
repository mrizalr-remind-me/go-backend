package model

import (
	"time"

	"github.com/google/uuid"
)

type base struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

func (b *base) Create() {
	b.ID = uuid.New()

	createTime := time.Now()
	b.CreatedAt = &createTime
}

func (b *base) Update() {
	updateTime := time.Now()
	b.UpdatedAt = &updateTime
}
