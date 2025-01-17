package domain

import "otzovik-back/internal/domain/models"

// Repository interfaces
type UserRepository interface {
    GetAllUsers() ([]models.User, error)
    CreateUser(user *models.User) error
}

type ReviewRepository interface {
    GetAllReviews() ([]models.Review, error)
    GetReviewById(id *int) (models.Review, error)
    CreateReview(review *models.Review) error
}

// Service interfaces
type UserService interface {
    GetAllUsers() ([]models.User, error)
    CreateUser(user *models.User) error
}

type ReviewService interface {
    GetAllReviews() ([]models.Review, error)
    GetReviewById(id *int) (models.Review, error)
    CreateReview(review *models.Review) error
} 