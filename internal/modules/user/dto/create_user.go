package dto

type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email,max=255" validate:"required,email,max=255"`
	Name  string `json:"name" binding:"required,min=2,max=100" validate:"required,min=2,max=100"`
}
