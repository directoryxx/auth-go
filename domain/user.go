package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        	uint           `gorm:"primaryKey"`
	Name		string		   `json:"name"`
	Username	string         `json:"username"`
	Password	string         `json:"password"`
	RoleID		uint		   `json:"role_id"`
	Role		Role		   `gorm:"foreignKey:RoleID"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}
