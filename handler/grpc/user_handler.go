package grpc

import (
	"context"

	"github.com/alvinfebriando/gin-gorm-skeleton/dto"
	pb "github.com/alvinfebriando/gin-gorm-skeleton/proto"
	"github.com/alvinfebriando/gin-gorm-skeleton/usecase"
)

type UserHandler struct {
	pb.UnimplementedUserServer
	uu usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *UserHandler {
	return &UserHandler{uu: uu}
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.TokenResponse, error) {
	request := dto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := request.Validate(); err != nil {
		return nil, err
	}

	user := request.ToUser()
	token, err := h.uu.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.TokenResponse{
		Token: token,
	}, nil
}

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserResponse, error) {
	request := dto.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := request.Validate(); err != nil {
		return nil, err
	}

	user := request.ToUser()
	createdUser, err := h.uu.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:    uint64(createdUser.Id),
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}, nil
}
