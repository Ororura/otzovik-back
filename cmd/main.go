package main

import (
	"log"
	"otzovik-back/api"
	"otzovik-back/config"
	"otzovik-back/repositories"
	"otzovik-back/services"
)

func main() {
	db := config.SetupDataBase()

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	reviewRepo := repositories.NewReviewRepo(db)
	reviewService := services.NewReviewService(reviewRepo)
	reviewHandler := api.NewReviewHandler(reviewService)

	router := api.MainRoute(userHandler, reviewHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
