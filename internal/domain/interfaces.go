package domain

import (
	"mime/multipart"
	"otzovik-back/internal/domain/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
}

type ReviewRepository interface {
	GetAllReviews() ([]models.Review, error)
	GetReviewById(id *int) (models.Review, error)
	CreateReview(review *models.Review) error
}

type ChatRepository interface {
	SaveMessage(message *models.Messages) error
	GetMessage(id *int) ([]*models.Messages, error)
}

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
}

type ChatService interface {
	SaveMessage(message *models.Messages) error
	GetMessage(id *int) ([]*models.Messages, error)
}

type ReviewService interface {
	GetAllReviews() ([]models.Review, error)
	GetReviewById(id *int) (models.Review, error)
	CreateReview(review *models.Review, image *multipart.FileHeader) error
	GetImage(imagePath string) (string, error)
}
