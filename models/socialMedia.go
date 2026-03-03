package models

import "time"

type SocialMedia struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	URL       string    `json:"url" gorm:"type:varchar(255);not null"`
	Image     string    `json:"image" gorm:"type:varchar(255)"`
	Status    bool      `json:"status" gorm:"default:true"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
