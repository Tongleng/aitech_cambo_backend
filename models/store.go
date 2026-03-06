package models

import "time"

type Store struct {
	ID          uint          `json:"id" gorm:"primaryKey"`
	Name        string        `json:"name" gorm:"type:varchar(255);not null"`
	Slug        string        `json:"slug" gorm:"type:varchar(255);uniqueIndex;not null;size:255"`
	Email       string        `json:"email" gorm:"type:varchar(100);unique;not null"`
	Phone       string        `json:"phone" gorm:"type:varchar(50)"`
	LogoURL     string        `json:"logoUrl" gorm:"type:varchar(255)"`
	BannerURL   string        `json:"bannerUrl" gorm:"type:varchar(255)"`
	Description string        `json:"description" gorm:"type:text"`
	WebsiteURL  string        `json:"websiteUrl" gorm:"type:varchar(255)"`
	TelegramURL string        `json:"telegramUrl" gorm:"type:varchar(255)"`
	FacebookURL string        `json:"facebookUrl" gorm:"type:varchar(255)"`
	Country     string        `json:"country" gorm:"type:varchar(100);default:'Cambodia'"`
	CategoryID  uint          `json:"categoryId"`
	Category    StoreCategory `json:"category" gorm:"foreignKey:CategoryID"`
	IsOpening   bool          `json:"isOpening" gorm:"default:true"`
	OpeningTime string        `json:"openingTime" gorm:"type:varchar(100)"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}
