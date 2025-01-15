package repositories

import (
	"gorm.io/gorm"
	"otzovik-back/model"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepo) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}
