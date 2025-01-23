package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"otzovik-back/config"
	"otzovik-back/internal/app/services"
	"otzovik-back/internal/platform/database/repositories"
	"otzovik-back/internal/platform/http/handlers"
	"otzovik-back/internal/platform/http/router"
)

func main() {
	cfg := config.LoadConfig()

	db := config.SetupDatabase(cfg)

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	userRepo := repositories.NewUserRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)

	userService := services.NewUserService(userRepo)
	reviewService := services.NewReviewService(reviewRepo, cfg.UploadDir)

	userHandler := handlers.NewUserHandler(userService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	websocketHandler := handlers.NewWebsocketHandler(upgrader)

	newRouter := router.NewRouter(userHandler, reviewHandler, websocketHandler)
	newRouter.SetupRoutes()

	if err := newRouter.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
