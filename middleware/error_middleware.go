package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
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

		message := strings.Split(err.Error(), "\n")
		switch {
		case err.Error() == "invalid request":
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: []string{"invalid request"},
			})
		case errors.Is(err, io.EOF):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: message,
			})
		case errors.As(err, &sErr):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: message,
			})
		case errors.As(err, &uErr):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: message,
			})
		case errors.As(err, &vErr):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: handleValidationError(vErr),
			})
		case isClientError:
			c.AbortWithStatusJSON(cErr.GetCode(), dto.Response{
				Error: message,
			})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
				Error: message,
			})
		}

	}
}

func handleValidationError(err validator.ValidationErrors) []string {
	output := make([]string, 0)
	for _, fieldError := range err {
		output = append(output, parseValidationError(fieldError))
	}
	return output
}

func parseValidationError(err validator.FieldError) string {
	field := strings.ToLower(err.Field())
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("invalid email value")
	case "alpha":
		return fmt.Sprintf("%s should consist of letters only", field)
	case "oneof":
		return fmt.Sprintf("%s's value should be one of [%v]", field, err.Param())
	case "startswith":
		return fmt.Sprintf("%s value should starts with %s", field, err.Param())
	case "numeric":
		return fmt.Sprintf("%s value should be numeric", field)
	case "len":
		return fmt.Sprintf("%s value length should be exactly %v", field, err.Param())
	}
	return err.Error()
}
