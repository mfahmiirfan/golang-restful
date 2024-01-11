package exception

type ForbiddenError struct {
	Message string
}

func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{
		Message: message,
	}
}

func (v *ForbiddenError) Error() string {
	return v.Message
}
