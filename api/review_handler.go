package api

import (
	"net/http"
	"otzovik-back/model"
	"otzovik-back/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	Service *services.ReviewService
}

func NewReviewHandler(service *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{Service: service}
}

func (h *ReviewHandler) GetReviews(c *gin.Context) {
	reviews, err := h.Service.GetAllReviews()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) GetReviewById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	review, err := h.Service.GetReviewById(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var review model.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}