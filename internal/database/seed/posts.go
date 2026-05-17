package seed

import (
	"belajargo/internal/models"
	"belajargo/internal/services"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PostsSeeder(db *gorm.DB) error {
	posts := []models.Posts{
		{
			ID:        uuid.MustParse("d756d290-5951-4410-b088-f0575e2d20a8"),
			Title:     "Post 1",
			Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
			UserID:    uuid.MustParse("acca69d2-5a1d-4cff-8ebb-3d4520471549"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, post := range posts {
		if err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(&post).Error; err != nil {
			services.Log(services.Logger{
				Name:    "POSTS SEEDER",
				Message: fmt.Sprintf("Failed to seed posts: %v", err),
			})
			return err
		}
	}

	services.Log(services.Logger{
		Name:    "SEEDER",
		Message: "Posts seeded successfully",
	})
	return nil
}
