package router

import (
	"github.com/gin-gonic/gin"
	"otzovik-back/internal/platform/http/handlers"
	"otzovik-back/internal/platform/http/middleware"
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
	// Middleware
	r.engine.Use(middleware.Logger())
	r.engine.Use(middleware.Recovery())

	// API routes
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
		}
	}
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
} 