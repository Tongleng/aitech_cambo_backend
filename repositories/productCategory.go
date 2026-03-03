package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type ProductCategoryRepository struct {
	DB *gorm.DB
}

func (r *ProductCategoryRepository) Create(category *models.ProductCategory) error {
	return r.DB.Create(category).Error
}

func (r *ProductCategoryRepository) GetProductCategory() ([]models.ProductCategory, error) {
	var categories []models.ProductCategory
	err := r.DB.Preload("Children").Where("parent_id IS NULL").Find(&categories).Error
	return categories, err
}

func (r *ProductCategoryRepository) GetBySlug(slug string) (models.ProductCategory, error) {
	var category models.ProductCategory
	err := r.DB.Preload("Children").Where("slug = ?", slug).First(&category).Error
	return category, err
}

func (r *SocialMediaRepository) Update(id uint, social *models.SocialMedia) error {
	return r.DB.Model(&models.SocialMedia{}).Where("id = ?", id).Updates(social).Error
}

func (r *SocialMediaRepository) GetByID(id uint) (models.SocialMedia, error) {
	var social models.SocialMedia
	err := r.DB.First(&social, id).Error
	return social, err
}

func (r *ProductCategoryRepository) GetAllPaginated(offset int, limit int) ([]models.ProductCategory, int64, error) {
	var categories []models.ProductCategory
	var total int64

	r.DB.Model(&models.ProductCategory{}).Count(&total)

	err := r.DB.Limit(limit).Offset(offset).Order("sort_order asc").Find(&categories).Error

	return categories, total, err
}

func (r *ProductCategoryRepository) Update(id uint, category *models.ProductCategory) error {
	return r.DB.Model(&models.ProductCategory{}).Where("id = ?", id).Updates(category).Error
}
