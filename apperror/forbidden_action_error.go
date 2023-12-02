package apperror

type ForbiddenActionError struct {
	resource string
}

func (e *ForbiddenActionError) Error() string {
	return "forbidden action"
}

func NewForbiddenActionError() error {
	return NewClientError(&ForbiddenActionError{}).Forbidden()
}
