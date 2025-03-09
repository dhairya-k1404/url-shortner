package database

import (
	"fmt"
	"log"
	"url-shortner/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitPostgresGORM initializes the PostgreSQL connection using GORM.
func InitPostgresGORM() {

	dbURL := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		conf.POSTGRES_HOST,
		conf.POSTGRES_USER,
		conf.POSTGRES_PASSWORD,
		conf.POSTGRES_DB,
		conf.POSTGRES_PORT,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL via GORM: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL via GORM")
}
