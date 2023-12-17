package main

import (
	"github.com/alvinfebriando/gin-gorm-skeleton/appjwt"
	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
	grpchandler "github.com/alvinfebriando/gin-gorm-skeleton/handler/grpc"
	"github.com/alvinfebriando/gin-gorm-skeleton/hash"
	"github.com/alvinfebriando/gin-gorm-skeleton/repository"
	grpcserver "github.com/alvinfebriando/gin-gorm-skeleton/server/grpc"
	"github.com/alvinfebriando/gin-gorm-skeleton/usecase"
)

func main() {
	applogger.SetLogrusLogger()

	db, err := repository.GetConnection()
	if err != nil {
		applogger.Log.Error(err)
	}

	userRepo := repository.NewUserRepository(db)

	newJwt := appjwt.NewJwt()
	newHasher := hash.NewHasher()

	userUsecase := usecase.NewUserUsecase(userRepo, newJwt, newHasher)

	userHandler := grpchandler.NewUserHandler(userUsecase)

	handlers := grpcserver.Handlers{
		User: userHandler,
	}

	s := grpcserver.New(handlers)
	grpcserver.StartWithGracefulShutdown(s)
}
