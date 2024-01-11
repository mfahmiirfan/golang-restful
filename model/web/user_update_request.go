package web

type UserUpdateRequest struct {
	ID              int    `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm"`
	Role            string `json:"role" validate:"required"`
	Verified        bool   `json:"verified"`
	// PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	// Photo           string `json:"photo"`
}
