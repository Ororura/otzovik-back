package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otzovik-back/services"
	"strconv"
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
