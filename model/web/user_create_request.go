package web

type UserCreateRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	Role            string `json:"role" validate:"required"`
	Verified        bool   `json:"verified"`
	// PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	// Photo           string `json:"photo"`
}
