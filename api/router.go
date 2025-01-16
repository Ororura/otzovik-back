package api

import (
	"github.com/gin-gonic/gin"
)

// MainRoute объединяет маршруты для пользователей и отзывов
func MainRoute(userHandler *UserHandler, reviewHandler *ReviewHandler) *gin.Engine {
	router := gin.Default()

	// Роуты для пользователей
	userGroup := router.Group("/users")
	{
		userGroup.GET("", userHandler.GetUsers)    // Получить список пользователей
		userGroup.POST("", userHandler.CreateUser) // Создать пользователя
	}

	// Роуты для отзывов
	reviewGroup := router.Group("/reviews")
	{
		reviewGroup.GET("/:id", reviewHandler.GetReviewById) // Получить отзыв по ID
	}

	return router
}
