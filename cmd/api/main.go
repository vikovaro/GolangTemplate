package main

import (
	"GolangTemplate/internal/modules/auth"
	"log"

	"github.com/gin-gonic/gin"

	_ "GolangTemplate/docs"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"GolangTemplate/internal/config"
	"GolangTemplate/internal/database"
	"GolangTemplate/internal/middleware"
	"GolangTemplate/internal/modules/user"
)

// @title Project API
// @version 1.0
// @description API template
// @BasePath /api
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	db, err := database.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db error: %v", err)
	}

	router := gin.New()

	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.Auth(cfg))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	api := router.Group("/api")

	auth.RegisterRoutes(api, db, cfg)
	user.RegisterRoutes(api, db)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
