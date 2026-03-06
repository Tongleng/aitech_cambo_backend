package services

import (
	"backend/models"
	"backend/repositories"
	"strings"
)

type StoreService struct {
	Repo repositories.StoreRepository
}

func (s *StoreService) CreateStore(p *models.Store) error {
	if p.Slug == "" {
		p.Slug = strings.ToLower(strings.ReplaceAll(p.Name, " ", "-"))
	}
	return s.Repo.Create(p)
}

func (s *StoreService) GetAllStores() ([]models.Store, error) {
	return s.Repo.GetAll()
}

func (s *StoreService) GetStores(page int, limit int) ([]models.Store, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	return s.Repo.GetAllPaginated(page, limit)
}

func (s *StoreService) GetStoreByID(id uint) (models.Store, error) {
	return s.Repo.GetByID(id)
}

func (s *StoreService) GetStoresByCategory(categoryID uint, page int, limit int) ([]models.Store, int64, error) {
	return s.Repo.GetByCategoryID(categoryID, page, limit)
}

func (s *StoreService) UpdateStore(id uint, storeData *models.Store) (models.Store, error) {
	err := s.Repo.Update(id, storeData)
	if err != nil {
		return models.Store{}, err
	}
	return s.Repo.GetByID(id)
}
