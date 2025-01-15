package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"otzovik-back/model"
)

func SetupDataBase() *gorm.DB {
	dsn := "host=localhost user=toor password=aiosghjdaogyiphdioUAHSDIUhasiud1erASD2q231d dbname=otzovik port=7432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err)
	}

	return db
}
