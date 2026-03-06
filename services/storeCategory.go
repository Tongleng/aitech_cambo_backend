package services

import (
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"strings"
)

type StoreCategoryService struct {
	Repo repositories.StoreCategoryRepository
}

func (s *StoreCategoryService) CreateStoreCategory(cat *models.StoreCategory) (*models.StoreCategory, error) {
	if cat.Slug == "" {
		cat.Slug = strings.ToLower(strings.ReplaceAll(cat.Name, " ", "-"))
	}

	err := s.Repo.Create(cat)
	return cat, err
}

func (s *StoreCategoryService) GetStoreCategory() ([]models.StoreCategory, error) {
	return s.Repo.GetStoreCategory()
}

func (s *StoreCategoryService) GetStoreCategoryPaginated(page int, pageSize int) (*utils.PaginationResponse, error) {
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

func (s *StoreCategoryService) UpdateStoreCategory(id uint, cat *models.StoreCategory) error {
	if cat.Name != "" && cat.Slug == "" {
		cat.Slug = strings.ToLower(strings.ReplaceAll(cat.Name, " ", "-"))
	}
	return s.Repo.Update(id, cat)
}
