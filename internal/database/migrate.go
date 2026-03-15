package database

import (
	"GolangTemplate/internal/modules/user/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error { return db.AutoMigrate(&model.User{}) }
