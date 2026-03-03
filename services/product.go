package services

import (
	"backend/models"
	"backend/repositories"
	"strings"
)

type ProductService struct {
	Repo repositories.ProductRepository
}

func (s *ProductService) CreateProduct(p *models.Product) error {
	if p.Slug == "" {
		p.Slug = strings.ToLower(strings.ReplaceAll(p.Name, " ", "-"))
	}
	return s.Repo.Create(p)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.Repo.GetAll()
}

func (s *ProductService) GetProducts(page int, limit int) ([]models.Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	return s.Repo.GetAllPaginated(page, limit)
}

func (s *ProductService) GetProductByID(id uint) (models.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *ProductService) GetProductsByCategory(categoryID uint, page int, limit int) ([]models.Product, int64, error) {
	return s.Repo.GetByCategoryID(categoryID, page, limit)
}
