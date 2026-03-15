package service

import (
	"GolangTemplate/internal/modules/user/model"
	"GolangTemplate/internal/modules/user/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(user *model.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetByUsername(username string) (*model.User, error) {
	return s.repo.GetByUsername(username)
}

func (s *UserService) Update(user *model.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}
