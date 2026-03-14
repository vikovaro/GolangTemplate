package auth

import (
	"GolangTemplate/internal/modules/auth/handler"
	"GolangTemplate/internal/modules/auth/service"
	"GolangTemplate/internal/modules/user/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"GolangTemplate/internal/config"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg)
	authHandler := handler.NewAuthHandler(authService)

	auth := rg.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
}
