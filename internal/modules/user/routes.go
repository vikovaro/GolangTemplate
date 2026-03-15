package user

import (
	"GolangTemplate/internal/modules/user/handler"
	"GolangTemplate/internal/modules/user/repository"
	"GolangTemplate/internal/modules/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	users := rg.Group("/users")
	{
		users.GET("/:id", h.GetByID)
		users.PUT("/:id", h.Update)
		users.DELETE("/:id", h.Delete)
	}
}
