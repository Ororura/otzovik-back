package repositories

import (
	"gorm.io/gorm"
	"otzovik-back/model"
)

type ReviewRepo struct {
	DB *gorm.DB
}

func NewReviewRepo(db *gorm.DB) *ReviewRepo {
	return &ReviewRepo{DB: db}
}

func (r *ReviewRepo) GetAllReviews() ([]model.Review, error) {
	var reviews []model.Review
	err := r.DB.Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepo) GetReviewById(id *int) (model.Review, error) {
	var review model.Review
	err := r.DB.First(&review, id).Error

	if err != nil {
		return model.Review{}, err
	}

	return review, nil
}

func (r *ReviewRepo) CreateReview(review *model.Review) error {
	return r.DB.Create(review).Error
}