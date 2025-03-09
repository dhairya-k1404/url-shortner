package url

import (
	"fmt"
	"net/http"
	"url-shortner/cache"
	"url-shortner/conf"
	urldto "url-shortner/dto/urlDto"
	"url-shortner/utils"

	"github.com/gin-gonic/gin"
)

func ShortenHandler(c *gin.Context) {
	var req urldto.URLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Generate a unique short code.
	shortCode := utils.GenerateShortCode()

	// Store the mapping using PostgreSQL and cache in Redis.
	if err := cache.SetMapping(shortCode, req.LongURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL mapping"})
		return
	}

	resp := urldto.URLResponse{
		ShortURL: fmt.Sprintf("%s%s", conf.BASE_URL, shortCode),
	}
	c.JSON(http.StatusOK, resp)
	return
}
