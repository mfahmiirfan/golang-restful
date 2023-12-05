package web

type UserResponse struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	// Role      string    `json:"role,omitempty"`
	// Photo     string    `json:"photo,omitempty"`
	// Provider  string    `json:"provider"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}
