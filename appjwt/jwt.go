package appjwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/alvinfebriando/gin-gorm-skeleton/config"
	"github.com/alvinfebriando/gin-gorm-skeleton/entity"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Id    uint
	Email string
}

type Jwt interface {
	GenerateToken(user *entity.User) (string, error)
	ValidateToken(tokenString string) (*entity.User, error)
}

type jwtImpl struct {
	secretKey []byte
}

func NewJwt() Jwt {
	return &jwtImpl{
		secretKey: []byte(config.New().JwtSecret),
	}
}

func (j *jwtImpl) GenerateToken(user *entity.User) (string, error) {
	userId := strconv.Itoa(int(user.Id))
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.New().JwtExpiryDuration * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.New().AppName,
			Subject:   userId,
		},
		Id: user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(j.secretKey)
	return signedString, err
}

func (j *jwtImpl) ValidateToken(tokenString string) (*entity.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		user := &entity.User{
			Id:    claims.Id,
			Email: claims.Email,
		}
		return user, nil
	}
	return nil, errors.New("invalid claims type")
}
