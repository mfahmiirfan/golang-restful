package web

import "github.com/google/uuid"

type UserUpdateRequest struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name" validate:"required"`
	Email    string    `json:"email" validate:"required"`
	Password string    `json:"password" validate:"required,min=8"`
	// PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	// Photo           string `json:"photo"`
}
