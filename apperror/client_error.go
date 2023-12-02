package apperror

import "net/http"

type ClientError struct {
	err  error
	code int
}

func (e *ClientError) Error() string {
	return e.err.Error()
}

func (e *ClientError) UnWrap() error {
	return e.err
}

func (e *ClientError) GetCode() int {
	return e.code
}

func (e *ClientError) BadRequest() error {
	e.code = http.StatusBadRequest
	return e
}
func (e *ClientError) Unauthorized() error {
	e.code = http.StatusUnauthorized
	return e
}
func (e *ClientError) Forbidden() error {
	e.code = http.StatusForbidden
	return e
}
func (e *ClientError) NotFound() error {
	e.code = http.StatusNotFound
	return e
}
func (e *ClientError) Conflict() error {
	e.code = http.StatusConflict
	return e
}
func NewClientError(err error, code ...int) *ClientError {
	statusCode := http.StatusBadRequest

	if len(code) > 0 {
		statusCode = code[0]
	}

	return &ClientError{
		err:  err,
		code: statusCode,
	}
}
