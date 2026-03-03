package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	FirstName string         `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string         `gorm:"type:varchar(100);not null" json:"last_name"`
	Email     string         `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string         `json:"password" gorm:"not null"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
