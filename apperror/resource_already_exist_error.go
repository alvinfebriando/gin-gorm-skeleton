package apperror

import (
	"fmt"
)

type ResourceAlreadyExist struct {
	resource string
	field    string
	value    any
}

func (e *ResourceAlreadyExist) Error() string {
	return fmt.Sprintf("%s with %s '%s' already exist", e.resource, e.field, e.value)
}

func NewResourceAlreadyExist(resource string, field string, value any) error {
	return NewClientError(&ResourceAlreadyExist{
		resource: resource,
		field:    field,
		value:    value,
	})
}
