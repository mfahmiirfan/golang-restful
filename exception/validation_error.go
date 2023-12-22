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

// func (e *ValidationError) Is(target error) bool {
// 	// Melakukan pengecekan apakah error merupakan instance dari CustomError
// 	_, ok := target.(*ValidationError)
// 	return ok
// }
