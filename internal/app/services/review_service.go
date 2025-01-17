package services

import (
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"
)

type ReviewService struct {
	repo domain.ReviewRepository
}

func NewReviewService(repo domain.ReviewRepository) domain.ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) GetAllReviews() ([]models.Review, error) {
	return s.repo.GetAllReviews()
}

func (s *ReviewService) GetReviewById(id *int) (models.Review, error) {
	return s.repo.GetReviewById(id)
}

func (s *ReviewService) CreateReview(review *models.Review) error {
	return s.repo.CreateReview(review)
} 