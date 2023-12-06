package exception

type LoginError struct {
	Message string
}

func NewLoginError(message string) LoginError {
	return LoginError{
		Message: message,
	}
}

func (v *LoginError) Error() string {
	return v.Message
}
