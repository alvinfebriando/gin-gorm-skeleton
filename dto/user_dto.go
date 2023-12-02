package dto

import "github.com/alvinfebriando/gin-gorm-skeleton/entity"

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (r *RegisterRequest) ToUser() *entity.User {
	return &entity.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
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
