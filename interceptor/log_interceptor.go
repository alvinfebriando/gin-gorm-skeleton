package interceptor

import (
	"context"
	"fmt"
	"time"

	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	startTime := time.Now()
	p, _ := peer.FromContext(ctx)
	m, err := handler(ctx, req)

	endTime := time.Now()

	durationInt := endTime.Sub(startTime)
	duration := durationInt.Microseconds()
	unit := "μs"
	if durationInt.Microseconds() > 1000 {
		duration = durationInt.Milliseconds()
		unit = "ms"
	}

	fields := map[string]any{
		"type":         "REQUEST",
		"network_type": p.Addr.Network(),
		"address":      p.Addr.String(),
		"uri":          info.FullMethod,
		"duration":     fmt.Sprintf("%d %s", duration, unit),
	}
	if err != nil {
		applogger.Log.WithFields(fields).Error(err)
		return m, err
	}
	applogger.Log.WithFields(fields).Info("request processed")

	return m, err
}
