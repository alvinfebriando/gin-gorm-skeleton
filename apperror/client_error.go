package apperror

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type ClientError struct {
	err      error
	httpCode int
	grpcCode codes.Code
}

func (e *ClientError) Error() string {
	return e.err.Error()
}

func (e *ClientError) UnWrap() error {
	return e.err
}

func (e *ClientError) HttpStatusCode() int {
	return e.httpCode
}

func (e *ClientError) GrpcStatusCode() codes.Code {
	return e.grpcCode
}

func (e *ClientError) BadRequest() error {
	e.httpCode = http.StatusBadRequest
	e.grpcCode = codes.InvalidArgument
	return e
}
func (e *ClientError) Unauthorized() error {
	e.httpCode = http.StatusUnauthorized
	e.grpcCode = codes.Unauthenticated
	return e
}
func (e *ClientError) Forbidden() error {
	e.httpCode = http.StatusForbidden
	e.grpcCode = codes.PermissionDenied
	return e
}
func (e *ClientError) NotFound() error {
	e.httpCode = http.StatusNotFound
	e.grpcCode = codes.NotFound
	return e
}
func (e *ClientError) Conflict() error {
	e.httpCode = http.StatusConflict
	e.grpcCode = codes.AlreadyExists
	return e
}
func NewClientError(err error) *ClientError {
	httpCode := http.StatusBadRequest
	grpcCode := codes.InvalidArgument

	return &ClientError{
		err:      err,
		httpCode: httpCode,
		grpcCode: grpcCode,
	}
}
