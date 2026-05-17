package dto

import (
	"belajargo/internal/config"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

type Router struct {
	DB  *gorm.DB
	API huma.API
	CFG *config.Config
}
