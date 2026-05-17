package seed

import (
	"belajargo/internal/models"
	"belajargo/internal/services"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UsersSeeder(db *gorm.DB) error {
	users := []models.Users{
		{
			ID:        uuid.MustParse("acca69d2-5a1d-4cff-8ebb-3d4520471549"),
			Name:      "Admin",
			Email:     "admin@gmail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range users {
		if err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "email"}},
			DoUpdates: clause.AssignmentColumns([]string{"id", "name", "updated_at"}),
		}).Create(&user).Error; err != nil {
			return err
		}
	}

	services.Log(services.Logger{
		Name:    "SEEDER",
		Message: "Users seeded successfully",
	})
	return nil
}
