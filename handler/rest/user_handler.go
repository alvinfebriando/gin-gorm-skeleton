package rest

import (
	"net/http"

	"github.com/alvinfebriando/gin-gorm-skeleton/dto"
	"github.com/alvinfebriando/gin-gorm-skeleton/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uu usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		uu: uu,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var request dto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(err)
		return
	}
	if err := request.Validate(); err != nil {
		_ = c.Error(err)
		return
	}

	user := request.ToUser()
	createdUser, err := h.uu.Register(c.Request.Context(), user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Data: dto.NewFromUser(createdUser),
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(err)
		return
	}
	if err := request.Validate(); err != nil {
		_ = c.Error(err)
		return
	}

	user := request.ToUser()
	token, err := h.uu.Login(c.Request.Context(), user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Data: token,
	})
}
