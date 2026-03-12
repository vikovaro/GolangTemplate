package handler

import (
	"GolangTemplate/internal/modules/user/service"
)

type UserHandler struct {
	service *service.UserService
}
