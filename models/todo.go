package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
