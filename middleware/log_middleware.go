package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		requestId := uuid.New()
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "request_id", requestId)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		endTime := time.Now()

		fields := map[string]any{
			"type":       "REQUEST",
			"request_id": requestId,
			"client_ip":  c.ClientIP(),
			"method":     c.Request.Method,
			"uri":        c.Request.RequestURI,
			"duration":   fmt.Sprintf("%d%s", endTime.Sub(startTime).Milliseconds(), " ms"),
			"status":     c.Writer.Status(),
		}
		applogger.Log.WithFields(fields).Info("request processed")
	}
}
