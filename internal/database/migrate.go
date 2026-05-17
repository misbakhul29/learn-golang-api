package database

import (
	"belajargo/internal/models"
	"belajargo/internal/services"
	"fmt"

	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Users{},
		&models.Posts{},
	)
	if err != nil {
		services.Log(services.Logger{
			Name:    "DATABASE",
			Message: fmt.Sprintf("Failed to migrate database: %v", err),
		})
		return err
	}

	services.Log(services.Logger{
		Name:    "DATABASE",
		Message: "Database migrated successfully",
	})
	return nil
}
