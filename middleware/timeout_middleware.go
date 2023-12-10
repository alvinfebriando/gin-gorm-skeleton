package middleware

import (
	"context"
	"time"

	"github.com/alvinfebriando/gin-gorm-skeleton/config"
	"github.com/gin-gonic/gin"
)

func Timeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, cancel := context.WithTimeout(ctx, config.New().RequestTimeout*time.Second)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
