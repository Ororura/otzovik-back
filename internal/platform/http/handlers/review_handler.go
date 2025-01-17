package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "otzovik-back/internal/domain"
    "otzovik-back/internal/domain/models"
    "strconv"
)

type ReviewHandler struct {
    service domain.ReviewService
}

func NewReviewHandler(service domain.ReviewService) *ReviewHandler {
    return &ReviewHandler{service: service}
}

func (h *ReviewHandler) GetReviews(c *gin.Context) {
    reviews, err := h.service.GetAllReviews()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) GetReviewById(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    review, err := h.service.GetReviewById(&id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
    var review models.Review
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.CreateReview(&review); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, review)
} 