package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	UserId    uuid.UUID // foreign key
	User      *User     //back-reference
	CreatedAt time.Time
	UpdatedAt time.Time
}
