package services

import (
	"backend/models"
	"backend/repositories"
)

type SocialMediaService struct {
	Repo repositories.SocialMediaRepository
}

func (s *SocialMediaService) Create(social *models.SocialMedia) error {
	return s.Repo.Create(social)
}

func (s *SocialMediaService) GetAllSocials() ([]models.SocialMedia, error) {
	return s.Repo.GetAll()
}
