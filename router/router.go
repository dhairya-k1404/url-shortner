package router

import (
	healthcheck "url-shortner/healthCheck"
	"url-shortner/services/redirect"
	"url-shortner/services/url"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns a Gin engine with all routes defined.
func SetupRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define routes
	r.POST("/shorten", url.ShortenHandler)
	r.GET("/:shortCode", redirect.RedirectHandler)

	r.GET("/health", healthcheck.HealthCheckHandler())

	return r
}
