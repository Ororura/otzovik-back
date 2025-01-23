package router

import (
	"otzovik-back/internal/platform/http/handlers"
	"otzovik-back/internal/platform/http/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine           *gin.Engine
	userHandler      *handlers.UserHandler
	reviewHandler    *handlers.ReviewHandler
	websocketHandler *handlers.WebsocketHandler
}

func NewRouter(userHandler *handlers.UserHandler, reviewHandler *handlers.ReviewHandler, websocketHandler *handlers.WebsocketHandler) *Router {
	return &Router{
		engine:           gin.Default(),
		userHandler:      userHandler,
		reviewHandler:    reviewHandler,
		websocketHandler: websocketHandler,
	}
}

func (r *Router) SetupRoutes() {
	r.engine.Use(middleware.Logger())
	r.engine.Use(middleware.Recovery())

	api := r.engine.Group("/api/v1")
	{
		websockets := api.Group("/ws")
		{
			websockets.GET("", r.websocketHandler.InitWebsocket) // Привязка обработчика WebSocket
		}

		users := api.Group("/users")
		{
			users.GET("", r.userHandler.GetUsers)
			users.POST("", r.userHandler.CreateUser)
		}

		reviews := api.Group("/reviews")
		{
			reviews.GET("", r.reviewHandler.GetReviews)
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
