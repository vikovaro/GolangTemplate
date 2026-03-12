package service

import (
	"GolangTemplate/internal/modules/user/repository"
)

type UserService struct {
	repo *repository.UserRepository
}
