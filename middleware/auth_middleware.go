package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/alvinfebriando/gin-gorm-skeleton/apperror"
	"github.com/alvinfebriando/gin-gorm-skeleton/appjwt"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		token, err := extractBearerToken(bearerToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		newJwt := appjwt.NewJwt()
		claims, err := newJwt.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": apperror.NewInvalidTokenError().Error(),
			})
			return
		}

		newContext := context.WithValue(c.Request.Context(), "user_id", claims.Id)
		c.Request = c.Request.WithContext(newContext)
		c.Next()
	}
}

func extractBearerToken(bearerToken string) (string, error) {
	if bearerToken == "" {
		return "", apperror.NewMissingTokenError()
	}
	token := strings.Split(bearerToken, " ")
	if len(token) != 2 {
		return "", apperror.NewInvalidTokenError()
	}
	return token[1], nil
}
