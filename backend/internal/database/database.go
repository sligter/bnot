package database

import (
	"bnot/backend/internal/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Note{}, &models.Version{})
}
