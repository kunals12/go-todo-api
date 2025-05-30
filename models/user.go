package models

import "time"

type User struct {
	Id        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string `json:"name"`
	Todos     []Todo `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
