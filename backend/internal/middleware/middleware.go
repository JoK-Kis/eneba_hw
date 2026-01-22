package middleware

import (
	"github.com/JoK-Kis/eneba-backend/config"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := config.GetFrontendOrigin()
		
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}