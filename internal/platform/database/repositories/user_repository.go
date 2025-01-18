package repositories

import (
	"gorm.io/gorm"
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}
