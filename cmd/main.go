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

	db, err := config.ConnectDB(cfg)

	if err != nil {
		log.Fatalf("Databse error: %v", err)
	}

	defer db.Close()

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	userRepo := repositories.NewUserRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)
	chatRepo := repositories.NewChatRepo(db)

	userService := services.NewUserService(userRepo)
	reviewService := services.NewReviewService(reviewRepo, cfg.UploadDir)
	chatService := services.NewChatService(chatRepo)

	userHandler := handlers.NewUserHandler(userService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	websocketHandler := handlers.NewWebsocketHandler(chatService, upgrader)

	newRouter := router.NewRouter(userHandler, reviewHandler, websocketHandler)
	newRouter.SetupRoutes()

	if err := newRouter.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
