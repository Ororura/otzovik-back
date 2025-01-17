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
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Инициализация базы данных
	db := config.SetupDatabase(cfg)

	// Инициализация репозиториев
	userRepo := repositories.NewUserRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)

	// Инициализация сервисов
	userService := services.NewUserService(userRepo)
	reviewService := services.NewReviewService(reviewRepo)

	// Инициализация обработчиков
	userHandler := handlers.NewUserHandler(userService)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	// Инициализация и настройка роутера
	router := router.NewRouter(userHandler, reviewHandler)
	router.SetupRoutes()

	// Запуск сервера
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
