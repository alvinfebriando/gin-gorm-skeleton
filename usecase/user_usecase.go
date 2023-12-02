package usecase

import (
	"context"

	"github.com/alvinfebriando/gin-gorm-skeleton/apperror"
	"github.com/alvinfebriando/gin-gorm-skeleton/appjwt"
	"github.com/alvinfebriando/gin-gorm-skeleton/entity"
	"github.com/alvinfebriando/gin-gorm-skeleton/hash"
	"github.com/alvinfebriando/gin-gorm-skeleton/repository"
	"github.com/alvinfebriando/gin-gorm-skeleton/valueobject"
)

type UserUsecase interface {
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
	Login(ctx context.Context, user *entity.User) (string, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
	jwt      appjwt.Jwt
	hash     hash.Hasher
}

func NewUserUsecase(
	userRepo repository.UserRepository,
	jwt appjwt.Jwt,
	hash hash.Hasher,
) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		jwt:      jwt,
		hash:     hash,
	}
}

func (u *userUsecase) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	emailQuery := valueobject.NewQuery().Condition("email", valueobject.Equal, user.Email)
	fetchedUser, err := u.userRepo.FindOne(ctx, emailQuery)
	if err != nil {
		return nil, err
	}
	if fetchedUser != nil {
		return nil, apperror.NewResourceAlreadyExist("user", "email", user.Email)
	}

	hashedPassword, err := u.hash.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	createdUser, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *userUsecase) Login(ctx context.Context, user *entity.User) (string, error) {
	query := valueobject.NewQuery().Condition("email", valueobject.Equal, user.Email)
	fetchedUser, err := u.userRepo.FindOne(ctx, query)
	if err != nil {
		return "", err
	}
	if fetchedUser == nil {
		return "", apperror.NewInvalidCredentialsError()
	}

	if !u.hash.Compare(fetchedUser.Password, user.Password) {
		return "", apperror.NewInvalidCredentialsError()
	}

	token, err := u.jwt.GenerateToken(fetchedUser)
	if err != nil {
		return "", err
	}

	return token, nil
}
