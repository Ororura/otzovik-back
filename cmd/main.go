package main

import (
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

	router := api.NewRouter(userHandler)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
