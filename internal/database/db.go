package database

import (
	"belajargo/internal/config"
	"belajargo/internal/services"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	cfg := config.LoadEnv()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.HostDB, cfg.PortDB, cfg.UserDB, cfg.PasswordDB, cfg.NameDB)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	services.Log(services.Logger{
		Name:    "DATABASE",
		Message: "Connected to database",
	})
	return db, nil
}
