package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"src/internal/models"
	"src/pkg/utils"
)

func CORSMiddleware() gin.HandlerFunc {
	utils.Logger.Info("setting up CORS middleware")
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Authorization")

		if c.Request.Method == http.MethodOptions {
			c.Status(http.StatusOK)
			return
		}
		c.Next()
	}
}

// AuthMiddleware checks for authentication of a user
func AuthMiddleware() func(ctx *gin.Context) {
	utils.Logger.Info("setting up Auth middleware")
	return func(c *gin.Context) {
		if _, err := models.ExtractAuthToken(c); err != nil {
			handleAuthError(c, models.Error{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}
		c.Next()
	}
}
