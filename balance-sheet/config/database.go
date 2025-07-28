package config

import (
	"balance-sheet/models"
	// "fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=Parthu732 dbname=balance-sheet port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	// Auto-migrate models
	err = DB.AutoMigrate(&models.User{}, &models.MorningRequest{}, &models.EveningRequest{})
	if err != nil {
		log.Fatal("Auto migration failed", err)
	}

}
