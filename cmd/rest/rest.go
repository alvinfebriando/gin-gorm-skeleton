package main

import (
	"github.com/alvinfebriando/gin-gorm-skeleton/appjwt"
	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
	resthandler "github.com/alvinfebriando/gin-gorm-skeleton/handler/rest"
	"github.com/alvinfebriando/gin-gorm-skeleton/hash"
	"github.com/alvinfebriando/gin-gorm-skeleton/repository"
	restrouter "github.com/alvinfebriando/gin-gorm-skeleton/router/rest"
	restserver "github.com/alvinfebriando/gin-gorm-skeleton/server/rest"
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
	userHandler := resthandler.NewUserHandler(userUsecase)

	handlers := restrouter.Handlers{
		User: userHandler,
	}

	r := restrouter.New(handlers)

	s := restserver.New(r)

	restserver.StartWithGracefulShutdown(s)
}
