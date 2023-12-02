package apperror

import (
	"fmt"
)

type ResourceStateError struct {
	resource string
	field    string
	state    string
}

func (e *ResourceStateError) Error() string {
	if e.field == "" {
		return fmt.Sprintf("%s %s", e.resource, e.state)
	}
	return fmt.Sprintf("%s %s %s", e.resource, e.field, e.state)
}

func NewResourceStateError(resource string, field string, state string) error {
	return NewClientError(&ResourceStateError{
		resource: resource,
		field:    field,
		state:    state,
	})
}
