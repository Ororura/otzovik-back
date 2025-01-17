package handlers

import (
	"net/http"
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"
	"strconv"

	"github.com/gin-gonic/gin"
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

    if err := c.ShouldBind(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    file, _ := c.FormFile("image")

    if err := h.service.CreateReview(&review, file); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, review)
}

func (h *ReviewHandler) GetImage(c *gin.Context) {
    imagePath := c.Param("imagePath")
    
    fullPath, err := h.service.GetImage(imagePath)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
        return
    }

    c.File(fullPath)
} 