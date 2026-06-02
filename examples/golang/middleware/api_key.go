package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.GetHeader("X-API-Key")

		if apiKey != os.Getenv("ATTENDANCE_API_KEY") {

			c.AbortWithStatusJSON(401, gin.H{
				"success": false,
				"message": "Invalid API Key",
			})

			return
		}

		c.Next()
	}
}