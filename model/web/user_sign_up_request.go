package web

type UserSignUpRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	Role            string `json:"role" validate:"required"`
	// Photo           string `json:"photo"`
}
