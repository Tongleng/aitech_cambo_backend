package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type StoreCategoryRepository struct {
	DB *gorm.DB
}

func (r *StoreCategoryRepository) Create(category *models.StoreCategory) error {
	return r.DB.Create(category).Error
}

func (r *StoreCategoryRepository) GetStoreCategory() ([]models.StoreCategory, error) {
	var categories []models.StoreCategory
	err := r.DB.Preload("Children").Where("parent_id IS NULL").Find(&categories).Error
	return categories, err
}

func (r *StoreCategoryRepository) GetBySlug(slug string) (models.StoreCategory, error) {
	var category models.StoreCategory
	err := r.DB.Preload("Children").Where("slug = ?", slug).First(&category).Error
	return category, err
}

func (r *StoreCategoryRepository) Update(id uint, category *models.StoreCategory) error {
	return r.DB.Model(&models.StoreCategory{}).Where("id = ?", id).Updates(category).Error
}

func (r *StoreCategoryRepository) GetByID(id uint) (models.StoreCategory, error) {
	var category models.StoreCategory
	err := r.DB.First(&category, id).Error
	return category, err
}

func (r *StoreCategoryRepository) GetAllPaginated(offset int, limit int) ([]models.StoreCategory, int64, error) {
	var categories []models.StoreCategory
	var total int64

	r.DB.Model(&models.StoreCategory{}).Count(&total)

	err := r.DB.Limit(limit).Offset(offset).Order("sort_order asc").Find(&categories).Error

	return categories, total, err
}
