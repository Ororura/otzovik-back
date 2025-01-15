package services

import (
	"otzovik-back/model"
	"otzovik-back/repositories"
)

type UserService struct {
	Repo *repositories.UserRepo
}

func NewUserService(repo *repositories.UserRepo) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Repo.CreateUser(user)
}
