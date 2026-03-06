package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Category").Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetBySlug(slug string) (models.Product, error) {
	var product models.Product
	err := r.DB.Preload("Category").Where("slug = ?", slug).First(&product).Error
	return product, err
}
func (r *ProductRepository) GetAllPaginated(page int, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	r.DB.Model(&models.Product{}).Count(&total)

	offset := (page - 1) * limit

	err := r.DB.Preload("Category").
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&products).Error

	return products, total, err
}
func (r *ProductRepository) GetByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.DB.Preload("Category").First(&product, id).Error
	return product, err
}

func (r *ProductRepository) GetByCategoryID(categoryID uint, page int, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	r.DB.Model(&models.Product{}).Where("category_id = ?", categoryID).Count(&total)

	offset := (page - 1) * limit
	err := r.DB.Preload("Category").
		Where("category_id = ?", categoryID).
		Offset(offset).Limit(limit).
		Find(&products).Error

	return products, total, err
}

func (r *ProductRepository) Update(id uint, updatedProduct *models.Product) error {
	return r.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updatedProduct).Error
}
