package services

import (
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"
)

type UserService struct {
    repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.repo.GetAllUsers()
}

func (s *UserService) CreateUser(user *models.User) error {
    return s.repo.CreateUser(user)
} 