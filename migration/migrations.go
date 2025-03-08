package migration

import (
	"log"
	"url-shortner/database"
	"url-shortner/models"
)

func AutoMigrate() {

	err := database.DB.AutoMigrate(&models.URLMapping{})
	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}
}
