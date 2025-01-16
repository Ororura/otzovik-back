package services

import (
	"otzovik-back/model"
	"otzovik-back/repositories"
)

type ReviewService struct {
	Repo *repositories.ReviewRepo
}

func NewReviewService(repo *repositories.ReviewRepo) *ReviewService {
	return &ReviewService{Repo: repo}
}

func (s *ReviewService) GetAllReviews() ([]model.Review, error) {
	return s.Repo.GetAllReviews()
}

func (s *ReviewService) GetReviewById(id *int) (model.Review, error) {
	return s.Repo.GetReviewById(id)
}
