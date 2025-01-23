package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"
	"otzovik-back/internal/platform/utils"
	"path/filepath"
)

type ReviewService struct {
	repo      domain.ReviewRepository
	uploadDir string
}

func NewReviewService(repo domain.ReviewRepository, uploadDir string) domain.ReviewService {
	return &ReviewService{
		repo:      repo,
		uploadDir: uploadDir,
	}
}

func (s *ReviewService) GetAllReviews() ([]models.Review, error) {
	return s.repo.GetAllReviews()
}

func (s *ReviewService) GetReviewById(id *int) (models.Review, error) {
	return s.repo.GetReviewById(id)
}

func (s *ReviewService) CreateReview(review *models.Review, image *multipart.FileHeader) error {
	if image != nil {
		fileName, err := utils.SaveFile(image, s.uploadDir)
		if err != nil {
			return err
		}
		review.ImagePath = fileName
	}
	return s.repo.CreateReview(review)
}

func (s *ReviewService) GetImage(imagePath string) (string, error) {
	fullPath := filepath.Join(s.uploadDir, imagePath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return "", fmt.Errorf("image not found")
	}
	return fullPath, nil
}
