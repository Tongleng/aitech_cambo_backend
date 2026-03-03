package services

import (
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"strings"
)

type ProductCategoryService struct {
	Repo repositories.ProductCategoryRepository
}

func (s *ProductCategoryService) CreateProductCategory(cat *models.ProductCategory) (*models.ProductCategory, error) {
	if cat.Slug == "" {
		cat.Slug = strings.ToLower(strings.ReplaceAll(cat.Name, " ", "-"))
	}

	err := s.Repo.Create(cat)
	return cat, err
}

func (s *ProductCategoryService) GetProductCategory() ([]models.ProductCategory, error) {
	return s.Repo.GetProductCategory()
}

func (s *SocialMediaService) UpdateSocial(id uint, social *models.SocialMedia) error {
	return s.Repo.Update(id, social)
}

func (s *ProductCategoryService) GetProductCategoryPaginated(page int, pageSize int) (*utils.PaginationResponse, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	categories, total, err := s.Repo.GetAllPaginated(offset, pageSize)
	if err != nil {
		return nil, err
	}

	lastPage := int((total + int64(pageSize) - 1) / int64(pageSize))

	return &utils.PaginationResponse{
		Data:     categories,
		Total:    total,
		Page:     page,
		LastPage: lastPage,
	}, nil
}

func (s *ProductCategoryService) UpdateProductCategory(id uint, cat *models.ProductCategory) error {
	if cat.Name != "" && cat.Slug == "" {
		cat.Slug = strings.ToLower(strings.ReplaceAll(cat.Name, " ", "-"))
	}
	return s.Repo.Update(id, cat)
}
