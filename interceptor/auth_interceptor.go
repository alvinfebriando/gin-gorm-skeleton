package interceptor

import (
	"context"
	"log"
	"strings"

	"github.com/alvinfebriando/gin-gorm-skeleton/apperror"
	"github.com/alvinfebriando/gin-gorm-skeleton/appjwt"
	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if !isProtected(info.FullMethod) {
		m, err := handler(ctx, req)
		if err != nil {
			log.Println(err)
		}
		return m, err
	}

	md, ok := metadata.FromIncomingContext(ctx)
	applogger.Log.Info(md)
	if !ok {
		return nil, apperror.NewMissingMetadataError()
	}

	bearerToken := md.Get("authorization")
	if len(bearerToken) == 0 {
		return nil, apperror.NewInvalidTokenError()
	}
	token := strings.TrimPrefix(bearerToken[0], "Bearer ")

	jwt := appjwt.NewJwt()
	claims, err := jwt.ValidateToken(token)
	if err != nil {
		return nil, apperror.NewInvalidTokenError()
	}

	ctx = context.WithValue(ctx, "user_id", claims.Id)
	m, err := handler(ctx, req)
	if err != nil {
		log.Println(err)
	}

	return m, err
}

var protection = map[string]bool{
	"/user.User/Login":    false,
	"/user.User/Register": false,
}

func isProtected(route string) bool {
	if val, ok := protection[route]; ok {
		return val
	}
	return false
}
