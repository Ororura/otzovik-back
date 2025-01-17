package router

import (
	"otzovik-back/internal/platform/http/handlers"
	"otzovik-back/internal/platform/http/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine      *gin.Engine
	userHandler  *handlers.UserHandler
	reviewHandler *handlers.ReviewHandler
}

func NewRouter(userHandler *handlers.UserHandler, reviewHandler *handlers.ReviewHandler) *Router {
	return &Router{
		engine:        gin.Default(),
		userHandler:   userHandler,
		reviewHandler: reviewHandler,
	}
}

func (r *Router) SetupRoutes() {
	r.engine.Use(middleware.Logger())
	r.engine.Use(middleware.Recovery())

	api := r.engine.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", r.userHandler.GetUsers)
			users.POST("", r.userHandler.CreateUser)
		}

		reviews := api.Group("/reviews")
		{
			reviews.GET("/:id", r.reviewHandler.GetReviewById)
			reviews.POST("", r.reviewHandler.CreateReview)
			reviews.GET("/images/:imagePath", r.reviewHandler.GetImage)
		}
	}

	r.engine.Static("/uploads", "./uploads")
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
} 