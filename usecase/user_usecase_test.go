package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/alvinfebriando/gin-gorm-skeleton/entity"
	"github.com/alvinfebriando/gin-gorm-skeleton/mocks"
	"github.com/alvinfebriando/gin-gorm-skeleton/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var user = &entity.User{
	Id:       1,
	Name:     "Alice",
	Email:    "alice@example.com",
	Password: "alice123",
}

var users = []*entity.User{user}

var generatedToken = []byte("TOKEN")

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

type UserUsecaseTestSuite struct {
	suite.Suite
	userRepo    *mocks.UserRepository
	jwt         *mocks.Jwt
	hash        *mocks.Hasher
	userUsecase usecase.UserUsecase
}

func (s *UserUsecaseTestSuite) SetupSubTest() {
	s.userRepo = mocks.NewUserRepository(s.T())
	s.jwt = mocks.NewJwt(s.T())
	s.hash = mocks.NewHasher(s.T())
	s.userUsecase = usecase.NewUserUsecase(s.userRepo, s.jwt, s.hash)
}

func (s *UserUsecaseTestSuite) TestUserUsecase_Register() {
	s.Run("should return new user", func() {
		s.userRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, nil)
		s.hash.On("Hash", mock.Anything).Return(generatedToken, nil)
		s.userRepo.On("Create", mock.Anything, mock.Anything).Return(user, nil)

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.userUsecase.Register(ctx, user)

		s.Equal(user, createdUser)
		s.NoError(err)
	})
	s.Run("should return error when there's an error when fetching user by email", func() {
		s.userRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.userUsecase.Register(ctx, user)

		s.Nil(createdUser)
		s.Error(err)
	})
	s.Run("should return error when user with same email already registered", func() {
		s.userRepo.On("FindOne", mock.Anything, mock.Anything).Return(user, nil)

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.userUsecase.Register(ctx, user)

		s.Nil(createdUser)
		s.Error(err)
	})
	s.Run("should return error when fail to hash the password", func() {
		s.userRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, nil)
		s.hash.On("Hash", mock.Anything).Return(generatedToken, errors.New(""))

		createdUser, err := s.userUsecase.Register(context.Background(), user)

		s.Nil(createdUser)
		s.Error(err)
	})
	s.Run("should return error when there's an error creating new user", func() {
		s.userRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, nil)
		s.hash.On("Hash", mock.Anything).Return(generatedToken, nil)
		s.userRepo.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.userUsecase.Register(ctx, user)

		s.Nil(createdUser)
		s.Error(err)
	})
}

func (s *UserUsecaseTestSuite) TestUserUsecase_Login() {
	s.Run("should return generatedToken", func() {
		s.userRepo.On("FindOne", mock.Anything, mock.Anything).Return(user, nil)
		s.hash.On("Compare", mock.Anything, mock.Anything).Return(true)
		s.jwt.On("GenerateToken", mock.Anything).Return("", nil)

		token, err := s.userUsecase.Login(context.Background(), user)

		s.NotNil(token)
		s.NoError(err)
	})
	s.Run("should return error when there's an error when searching for user", func() {
		s.userRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		token, err := s.userUsecase.Login(context.Background(), user)

		s.Equal("", token)
		s.Error(err)
	})
	s.Run("should return error when there's no user", func() {
		s.userRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, nil)

		token, err := s.userUsecase.Login(context.Background(), user)

		s.Equal("", token)
		s.Error(err)
	})
	s.Run("should return error when password is wrong", func() {
		s.userRepo.On("FindOne", mock.Anything, mock.Anything).Return(user, nil)
		s.hash.On("Compare", mock.Anything, mock.Anything).Return(false)

		token, err := s.userUsecase.Login(context.Background(), user)

		s.Equal("", token)
		s.Error(err)
	})
	s.Run("should return error when there's an error when generating generatedToken", func() {
		s.userRepo.On("FindOne", mock.Anything, mock.Anything).Return(user, nil)
		s.hash.On("Compare", mock.Anything, mock.Anything).Return(true)
		s.jwt.On("GenerateToken", mock.Anything).Return("", errors.New(""))

		token, err := s.userUsecase.Login(context.Background(), user)

		s.Equal("", token)
		s.Error(err)
	})
}
