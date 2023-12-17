package rest

import (
	"net/http"

	resthandler "github.com/alvinfebriando/gin-gorm-skeleton/handler/rest"
	"github.com/alvinfebriando/gin-gorm-skeleton/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User *resthandler.UserHandler
}

func New(handlers Handlers) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Timeout())
	router.Use(middleware.Logger())
	router.Use(middleware.Error())

	router.POST("/register", handlers.User.Register)
	router.POST("/login", handlers.User.Login)

	router.Use(middleware.Authentication())

	return router
}
