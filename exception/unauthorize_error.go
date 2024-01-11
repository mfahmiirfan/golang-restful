package exception

type UnauthorizeError struct {
	Message string
}

func NewUnauthorizeError(message string) *UnauthorizeError {
	return &UnauthorizeError{
		Message: message,
	}
}

func (v *UnauthorizeError) Error() string {
	return v.Message
}
