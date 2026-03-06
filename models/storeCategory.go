package models

import (
	"time"
)

type StoreCategory struct {
	ID          uint            `json:"id" gorm:"primaryKey"`
	Name        string          `json:"name" gorm:"type:varchar(255);not null"`
	Slug        string          `json:"slug" gorm:"type:varchar(255);uniqueIndex;not null;size:255"`
	Description string          `json:"description"`
	ImageURL    string          `json:"imageUrl"`
	ParentID    *uint           `json:"parentId" gorm:"index"`
	Parent      *StoreCategory  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children    []StoreCategory `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	SortOrder   int             `json:"sortOrder" gorm:"default:0"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}
