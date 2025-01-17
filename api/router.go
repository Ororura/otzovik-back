package api

import (
	"github.com/gin-gonic/gin"
)

func MainRoute(userHandler *UserHandler, reviewHandler *ReviewHandler) *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/users")
	{
		userGroup.GET("", userHandler.GetUsers)
		userGroup.POST("", userHandler.CreateUser)
	}

	reviewGroup := router.Group("/reviews")
	{
		reviewGroup.GET("/:id", reviewHandler.GetReviewById)
		reviewGroup.POST("", reviewHandler.CreateReview)
	}

	return router
}

