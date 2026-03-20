package main

import (
	"GolangTemplate/internal/config"
	"GolangTemplate/internal/database"
	"GolangTemplate/internal/middleware"
	"GolangTemplate/internal/modules/auth"
	"GolangTemplate/internal/modules/user"
	"log"

	_ "GolangTemplate/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title Project API
// @version 1.0
// @description API template
// @BasePath /api
func main() {
	// 1. Инициализация конфигурации
	cfg := initConfig()

	// 2. Подключение к базе данных
	db := initDatabase(cfg)

	// 3. Настройка роутера
	router := setupRouter(cfg, db)

	// 4. Запуск сервера
	startServer(router, cfg)
}

func initConfig() *config.Config {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Println("Config loaded successfully")
	return cfg
}

func initDatabase(cfg *config.Config) *gorm.DB {
	db, err := database.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")

	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")

	return db
}

func setupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Printf("Warning: failed to set trusted proxies: %v", err)
	}

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := router.Group("/api")

	// Public routes
	auth.RegisterRoutes(api, db, cfg)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.Auth(cfg))
	user.RegisterRoutes(protected, db)

	log.Println("Router configured successfully")
	return router
}

func startServer(router *gin.Engine, cfg *config.Config) {
	addr := ":" + cfg.Port
	log.Printf("Server starting on %s", addr)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
