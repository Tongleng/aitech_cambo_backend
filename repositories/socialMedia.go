package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type SocialMediaRepository struct {
	DB *gorm.DB
}

func (r *SocialMediaRepository) Create(social *models.SocialMedia) error {
	return r.DB.Create(social).Error
}

func (r *SocialMediaRepository) GetAll() ([]models.SocialMedia, error) {
	var socials []models.SocialMedia
	err := r.DB.Find(&socials).Error
	return socials, err
}

func (r *SocialMediaRepository) Delete(id uint) error {
	return r.DB.Delete(&models.SocialMedia{}, id).Error
}
