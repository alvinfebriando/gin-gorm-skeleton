package apperror

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	err error
}

func NewValidationError(err error) error {
	return NewClientError(&ValidationError{err})
}

func (e *ValidationError) Error() string {
	var vErr validator.ValidationErrors
	if errors.As(e.err, &vErr) {
		message := handleValidationError(vErr)
		return strings.Join(message, "\n")
	}
	return ""
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
