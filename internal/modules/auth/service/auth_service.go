package service

import (
	"GolangTemplate/internal/config"
	"GolangTemplate/internal/modules/user/repository"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: cfg.JWTSecret,
	}
}
