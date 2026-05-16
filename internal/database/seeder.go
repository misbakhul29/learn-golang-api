package database

import (
	"belajargo/internal/models"
	"belajargo/internal/services"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seeders(db *gorm.DB) error {
	users := []models.Users{
		{
			ID:        uuid.New(),
			Name:      "Admin",
			Email:     "admin@gmail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			services.Log(services.Logger{
				Name:    "DATABASE",
				Message: fmt.Sprintf("Failed to seed users: %v", err),
			})
			return err
		}
	}

	return nil
}
