package services

import (
	"backend/models"
	"backend/repositories"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo repositories.UserRepository
}

func (s *UserService) GetUsers() ([]models.User, error) {
	users, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) Register(user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	if err := s.Repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(email, password string) (*models.User, string, error) {
	secretKey := os.Getenv("DB_HOST")
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, "", err
	}

	return &user, t, nil
}
