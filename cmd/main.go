package main

import (
	"log"
	"otzovik-back/config"
	"otzovik-back/internal/app/services"
	"otzovik-back/internal/platform/database/repositories"
	"otzovik-back/internal/platform/http/handlers"
	"otzovik-back/internal/platform/http/router"
)

func main() {
	cfg := config.LoadConfig()

	db := config.SetupDatabase(cfg)

	userRepo := repositories.NewUserRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)

	userService := services.NewUserService(userRepo)
	reviewService := services.NewReviewService(reviewRepo, cfg.UploadDir)

	userHandler := handlers.NewUserHandler(userService)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	router := router.NewRouter(userHandler, reviewHandler)
	router.SetupRoutes()

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
