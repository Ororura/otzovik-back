package repositories

import (
	"database/sql"
	"fmt"
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) domain.ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) GetAllReviews() ([]models.Review, error) {
	query := `
			SELECT r.id, r.title, r.content, r.rating, r.image_path, u.id, u.username, u.password, c.id, c.name FROM reviews r 
			JOIN users u ON r.user_id = u.id
			JOIN categories c ON r.category_id = c.id`
	rows, err := r.db.Query(query)

	defer rows.Close()

	var reviews []models.Review

	for rows.Next() {
		var review models.Review
		if err := rows.Scan(
			&review.ID, &review.Title, &review.Content, &review.Rating, &review.ImagePath,
			&review.UserID, &review.UserName,
			&review.CategoryID, &review.CategoryName,
		); err != nil {
			return nil, fmt.Errorf("scan error %v", err)
		}

		reviews = append(reviews, review)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %v", err)
	}

	return reviews, nil
}

func (r *ReviewRepository) GetReviewById(id *int) (models.Review, error) {
	query := `
			SELECT r.id, r.title, r.content, r.rating, r.image_path, u.id, u.username, u.password, c.id, c.name FROM reviews r 
			JOIN users u ON r.user_id = u.id
			JOIN categories c ON r.category_id = c.id
			WHERE r.id = $1`

	var review models.Review
	var user models.User
	var category models.Category

	err := r.db.QueryRow(query, id).Scan(
		&review.ID, &review.Title, &review.Content, &review.Rating, &review.ImagePath,
		&user.ID, &user.Username, &user.Password,
		&category.ID, &category.Name,
	)
	if err != nil {
		return review, fmt.Errorf("ошибка сканирования данных: %v", err)
	}

	// Присваиваем данные в review
	review.Review = &user
	review.Category = &category

	return review, err
}

func (r *ReviewRepository) CreateReview(review *models.Review) error {
	query := "INSERT INTO reviews(title, content, rating, image_path, user_id, category_id) VALUES ($1, $2, $3, $4, $5, $6)"
	err := r.db.QueryRow(query, review.Title, review.Content, review.ImagePath, review.UserID, review.CategoryID).Scan(&review.ID)
	if err != nil {
		return fmt.Errorf("Ошибка при сохранении отзыва %v", err)
	}
	return nil
}
