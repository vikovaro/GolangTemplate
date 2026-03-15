package dto

type UpdateUserRequest struct {
	Email    *string `json:"email" binding:"omitempty,email,max=255" validate:"omitempty,email,max=255"`
	Username *string `json:"username" binding:"omitempty,min=2,max=100" validate:"omitempty,min=2,max=100"`
	Phone    *string `json:"phone" binding:"omitempty,min=2,max=30" validate:"omitempty,min=2,max=30"`
	Password *string `json:"password" binding:"omitempty,min=6,max=100" validate:"omitempty,min=6,max=100"`
}
