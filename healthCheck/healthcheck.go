package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "All systems operational 🚀",
		})
	}
}
