package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/alvinfebriando/gin-gorm-skeleton/apperror"
	"github.com/alvinfebriando/gin-gorm-skeleton/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0].Err

		var sErr *json.SyntaxError
		var uErr *json.UnmarshalTypeError
		var vErr validator.ValidationErrors
		var cErr *apperror.ClientError

		isClientError := false
		if errors.As(err, &cErr) {
			isClientError = true
			err = cErr.UnWrap()
		}

		switch {
		case err.Error() == "invalid request":
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: errors.New("invalid request").Error(),
			})
		case errors.Is(err, io.EOF):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: err.Error(),
			})
		case errors.As(err, &sErr):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: sErr.Error(),
			})
		case errors.As(err, &uErr):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: uErr.Error(),
			})
		case errors.As(err, &vErr):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: strings.Split(vErr.Error(), "\n"),
			})
		case isClientError:
			c.AbortWithStatusJSON(cErr.GetCode(), dto.Response{
				Error: cErr.Error(),
			})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
				Error: err.Error(),
			})
		}

	}
}
