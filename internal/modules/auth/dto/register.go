package dto

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone" validate:"required, phone"`
	Email    string `json:"email" validate:"required,email"`
}
