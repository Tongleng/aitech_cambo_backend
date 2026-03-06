package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type StoreRepository struct {
	DB *gorm.DB
}

func (r *StoreRepository) Create(store *models.Store) error {
	return r.DB.Create(store).Error
}

func (r *StoreRepository) GetAll() ([]models.Store, error) {
	var stores []models.Store
	err := r.DB.Preload("Category").Find(&stores).Error
	return stores, err
}

func (r *StoreRepository) GetBySlug(slug string) (models.Store, error) {
	var store models.Store
	err := r.DB.Preload("Category").Where("slug = ?", slug).First(&store).Error
	return store, err
}
func (r *StoreRepository) GetAllPaginated(page int, limit int) ([]models.Store, int64, error) {
	var stores []models.Store
	var total int64

	r.DB.Model(&models.Store{}).Count(&total)

	offset := (page - 1) * limit

	err := r.DB.Preload("Category").
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&stores).Error

	return stores, total, err
}
func (r *StoreRepository) GetByID(id uint) (models.Store, error) {
	var store models.Store
	err := r.DB.Preload("Category").First(&store, id).Error
	return store, err
}

func (r *StoreRepository) GetByCategoryID(categoryID uint, page int, limit int) ([]models.Store, int64, error) {
	var stores []models.Store
	var total int64

	r.DB.Model(&models.Store{}).Where("category_id = ?", categoryID).Count(&total)

	offset := (page - 1) * limit
	err := r.DB.Preload("Category").
		Where("category_id = ?", categoryID).
		Offset(offset).Limit(limit).
		Find(&stores).Error

	return stores, total, err
}

func (r *StoreRepository) Update(id uint, updatedStore *models.Store) error {
	return r.DB.Model(&models.Store{}).Where("id = ?", id).Updates(updatedStore).Error
}
