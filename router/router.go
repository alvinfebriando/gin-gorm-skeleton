package router

import (
	"net/http"

	"github.com/alvinfebriando/gin-gorm-skeleton/handler"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User *handler.UserHandler
}

func New(handlers Handlers) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/register", handlers.User.Register)
	router.POST("/login", handlers.User.Login)

	return router
}