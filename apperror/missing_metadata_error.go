package apperror

import (
	"errors"
)

var ErrMissingMetadata = errors.New("missing metadata")

func NewMissingMetadataError() error {
	return NewClientError(ErrMissingToken).BadRequest()
}
