package redirect

import (
	"net/http"
	"url-shortner/cache"

	"github.com/gin-gonic/gin"
)

func RedirectHandler(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if longURL, ok := cache.GetMapping(shortCode); ok {
		c.Redirect(http.StatusFound, longURL)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
}
