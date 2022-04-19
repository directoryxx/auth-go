package domain

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID        	uint           `gorm:"primaryKey"`
	Name		string		   `json:"name"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}

