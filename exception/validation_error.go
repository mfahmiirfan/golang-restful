package exception

type ValidationError struct {
	Message string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		Message: message,
	}
}

func (v *ValidationError) Error() string {
	return v.Message
}
