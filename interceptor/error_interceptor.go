package interceptor

import (
	"context"
	"errors"

	"github.com/alvinfebriando/gin-gorm-skeleton/apperror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	m, err := handler(ctx, req)

	if err == nil {
		return m, err
	}

	var cErr *apperror.ClientError

	isClientError := false

	if errors.As(err, &cErr) {
		isClientError = true
		err = cErr.UnWrap()
	}

	switch {
	case isClientError:
		return m, status.Error(cErr.GrpcStatusCode(), err.Error())
	case errors.Is(err, context.DeadlineExceeded):
		return m, status.Error(codes.DeadlineExceeded, err.Error())
	default:
		return m, status.Error(codes.Internal, err.Error())
	}
}
