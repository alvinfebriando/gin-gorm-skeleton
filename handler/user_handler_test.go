package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvinfebriando/gin-gorm-skeleton/apperror"
	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
	"github.com/alvinfebriando/gin-gorm-skeleton/dto"
	"github.com/alvinfebriando/gin-gorm-skeleton/entity"
	"github.com/alvinfebriando/gin-gorm-skeleton/handler"
	"github.com/alvinfebriando/gin-gorm-skeleton/mocks"
	"github.com/alvinfebriando/gin-gorm-skeleton/router"
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

var generatedToken = "TOKEN"

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

type UserHandlerTestSuite struct {
	suite.Suite
	router http.Handler
	uu     *mocks.UserUsecase
	uh     *handler.UserHandler
	rec    *httptest.ResponseRecorder
}

func (s *UserHandlerTestSuite) SetupSubTest() {
	applogger.SetLogrusLogger()
	s.uu = mocks.NewUserUsecase(s.T())
	s.uh = handler.NewUserHandler(s.uu)
	h := router.Handlers{
		User: s.uh,
	}
	s.router = router.New(h)
	s.rec = httptest.NewRecorder()
}

func (s *UserHandlerTestSuite) TestRegister() {
	s.Run("should return 201", func() {
		request := dto.RegisterRequest{
			Name:     "Alice",
			Email:    "alice@example.com",
			Password: "alice123",
		}
		s.uu.On("Register", mock.Anything, mock.Anything).Return(user, nil)
		response := dto.Response{
			Data: dto.NewFromUser(user),
		}

		req, _ := http.NewRequest(http.MethodPost, "/register", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusCreated, s.rec.Code)
		s.Equal(marshal(response), getBody(s.rec))
	})

	s.Run("should return 400 when request body is unexpected", func() {
		req, _ := http.NewRequest(http.MethodPost, "/register", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})

	s.Run("should return 400 when email already exists", func() {
		request := dto.RegisterRequest{
			Name:     "Alice",
			Email:    "alice@example.com",
			Password: "alice123",
		}
		err := apperror.NewResourceAlreadyExist("user", "email", request.Email)
		s.uu.On("Register", mock.Anything, mock.Anything).Return(nil, err)

		req, _ := http.NewRequest(http.MethodPost, "/register", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
}

func (s *UserHandlerTestSuite) TestLogin() {
	s.Run("should return 200", func() {
		request := dto.LoginRequest{
			Email:    "alice@example.com",
			Password: "alice123",
		}
		s.uu.On("Login", mock.Anything, mock.Anything).Return(generatedToken, nil)
		response := dto.Response{
			Data: generatedToken,
		}

		req, _ := http.NewRequest(http.MethodPost, "/login", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusOK, s.rec.Code)
		s.Equal(marshal(response), getBody(s.rec))
	})
	s.Run("should return 400 when request body is unexpected", func() {
		request := dto.LoginRequest{
			Email: "alice@example.com",
		}

		req, _ := http.NewRequest(http.MethodPost, "/login", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
	s.Run("should return 401", func() {
		request := dto.LoginRequest{
			Email:    "alice@example.com",
			Password: "wrong password",
		}
		s.uu.On("Login", mock.Anything, mock.Anything).Return("", apperror.NewInvalidCredentialsError())

		req, _ := http.NewRequest(http.MethodPost, "/login", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusUnauthorized, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
}
