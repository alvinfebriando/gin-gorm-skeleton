package dto

import (
	"strings"

	"github.com/alvinfebriando/gin-gorm-skeleton/apperror"
	"github.com/alvinfebriando/gin-gorm-skeleton/entity"
	"github.com/alvinfebriando/gin-gorm-skeleton/validator"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (r *RegisterRequest) Validate() error {
	password := strings.Trim(r.Password, " ")
	r.Password = password
	err := validator.New().Struct(r)
	if err != nil {
		return apperror.NewValidationError(err)
	}

	return nil
}

func (r *RegisterRequest) ToUser() *entity.User {
	return &entity.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginRequest) Validate() error {
	password := strings.Trim(r.Password, "")
	r.Password = password
	err := validator.New().Struct(r)
	if err != nil {
		return apperror.NewValidationError(err)
	}

	return nil
}

func (r *LoginRequest) ToUser() *entity.User {
	return &entity.User{
		Email:    r.Email,
		Password: r.Password,
	}
}

type UserResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewFromUser(user *entity.User) *UserResponse {
	return &UserResponse{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}
}
