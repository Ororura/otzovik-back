package repositories

import (
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) domain.ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Preload("User").Preload("Category").Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepository) GetReviewById(id *int) (models.Review, error) {
	var review models.Review
	err := r.db.First(&review, id).Error
	return review, err
}

func (r *ReviewRepository) CreateReview(review *models.Review) error {
	return r.db.Create(review).Error
}
