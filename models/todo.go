package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	UserId    uuid.UUID `json:"userId"`
	User      *User     `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
