package repositories

import (
	"database/sql"
	"fmt"
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := "SELECT id, username, email, password FROM users"
	rows, err := r.db.Query(query)
	var users []models.User

	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %v", err)
	}

	return users, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	err := r.db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("ошибка при создании пользователя: %v", err)
	}

	return nil
}
