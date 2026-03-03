package models

import "time"

type Product struct {
	ID          uint            `json:"id" gorm:"primaryKey"`
	Name        string          `json:"name" gorm:"type:varchar(255);not null"`
	Slug        string          `json:"slug" gorm:"type:varchar(255);uniqueIndex;not null;size:255"`
	Description string          `json:"description" gorm:"type:text"`
	Price       float64         `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock       int             `json:"stock" gorm:"default:0"`
	ImageURL    string          `json:"imageUrl"`
	CategoryID  uint            `json:"categoryId"`
	Category    ProductCategory `json:"category" gorm:"foreignKey:CategoryID"`
	Status      bool            `json:"status" gorm:"default:true"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}
