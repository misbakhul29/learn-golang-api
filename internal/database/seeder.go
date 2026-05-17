package database

import (
	"belajargo/internal/database/seed"

	"belajargo/internal/services"
	"fmt"

	"gorm.io/gorm"
)

func Seeders(db *gorm.DB) error {
	if err := seed.UsersSeeder(db); err != nil {
		services.Log(services.Logger{
			Name:    "SEEDER",
			Message: fmt.Sprintf("Failed to seed users: %v", err),
		})
		return err
	}

	if err := seed.PostsSeeder(db); err != nil {
		services.Log(services.Logger{
			Name:    "SEEDER",
			Message: fmt.Sprintf("Failed to seed posts: %v", err),
		})
		return err
	}

	return nil
}
