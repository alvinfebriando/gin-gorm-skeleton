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
		durationInt := endTime.Sub(startTime)
		duration := durationInt.Microseconds()
		unit := "μs"
		if durationInt.Microseconds() > 1000 {
			duration = durationInt.Milliseconds()
			unit = "ms"
		}

		fields := map[string]any{
			"type":       "REQUEST",
			"request_id": requestId,
			"client_ip":  c.ClientIP(),
			"method":     c.Request.Method,
			"uri":        c.Request.RequestURI,
			"duration":   fmt.Sprintf("%d%s", duration, unit),
			"status":     c.Writer.Status(),
		}
		err := c.Errors.Last()
		if err != nil {
			applogger.Log.WithFields(fields).Error(err)
			return
		}
		applogger.Log.WithFields(fields).Info("request processed")
	}
}
