package middleware

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Timeout() gin.HandlerFunc {
	const defaultTimeout = 5
	seconds, err := strconv.Atoi(os.Getenv("REQUEST_TIMEOUT"))
	if err != nil {
		seconds = defaultTimeout
	}

	duration := time.Duration(seconds) * time.Second
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, cancel := context.WithTimeout(ctx, duration)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
