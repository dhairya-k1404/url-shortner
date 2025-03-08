package main

import (
	"log"
	"url-shortner/cache"
	"url-shortner/database"
	"url-shortner/migration"
	"url-shortner/router"
)

func main() {

	// Get the router from the router package
	database.InitPostgresGORM()
	migration.AutoMigrate()
	cache.InitRedis()

	r := router.SetupRouter()

	log.Println("Server running on :8080")
	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
