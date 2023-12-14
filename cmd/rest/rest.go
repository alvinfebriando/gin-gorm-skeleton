package main

import (
	"github.com/alvinfebriando/gin-gorm-skeleton/appjwt"
	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
	"github.com/alvinfebriando/gin-gorm-skeleton/handler"
	"github.com/alvinfebriando/gin-gorm-skeleton/hash"
	"github.com/alvinfebriando/gin-gorm-skeleton/repository"
	"github.com/alvinfebriando/gin-gorm-skeleton/router"
	"github.com/alvinfebriando/gin-gorm-skeleton/server"
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
	userHandler := handler.NewUserHandler(userUsecase)

	handlers := router.Handlers{
		User: userHandler,
	}

	r := router.New(handlers)

	s := server.New(r)

	server.StartWithGracefulShutdown(s)
}
