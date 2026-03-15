package main

import (
	"GolangTemplate/internal/modules/auth"
	"log"

	"github.com/gin-gonic/gin"

	_ "GolangTemplate/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	if err := database.Migrate(db); err != nil {
		log.Fatalf("migration error: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.SetTrustedProxies(nil)

	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	auth.RegisterRoutes(api, db, cfg)

	protected := api.Group("/")
	protected.Use(middleware.Auth(cfg))
	{
		user.RegisterRoutes(protected, db)
	}

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
