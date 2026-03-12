package dto

type UpdateUserRequest struct {
	Email *string `json:"email" binding:"omitempty,email,max=255" validate:"omitempty,email,max=255"`
	Name  *string `json:"name" binding:"omitempty,min=2,max=100" validate:"omitempty,min=2,max=100"`
}
