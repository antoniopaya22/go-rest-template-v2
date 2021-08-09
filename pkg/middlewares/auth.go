package middlewares

import (
	"ai4am.com/go-restapi/pkg/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if !crypto.ValidateToken(authorizationHeader) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		} else {
			c.Next()
		}
	}
}
