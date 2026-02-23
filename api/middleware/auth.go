package middleware

import (
	"net/http"
	"strings"

	"github.com/Samir-Minddeft/go-backend-boilerplate/utils/helper"
	"github.com/gin-gonic/gin"
)

func AuthRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		jwt := strings.TrimPrefix(token, "Bearer ")

		claims, err := helper.VerifyJwtToken(jwt)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		switch role {
		case "admin":
			if claims["role"] != "admin" {
				c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
				c.Abort()
				return
			}
		case "user":
			if claims["role"] != "user" {
				c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
				c.Abort()
				return
			}
		default:
			c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
