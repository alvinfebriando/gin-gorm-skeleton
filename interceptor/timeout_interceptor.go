package interceptor

import (
	"context"
	"time"

	"github.com/alvinfebriando/gin-gorm-skeleton/config"
	"google.golang.org/grpc"
)

func TimeoutInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	ctx, cancel := context.WithTimeout(ctx, config.New().RequestTimeout*time.Second)
	defer cancel()

	return handler(ctx, req)
}
