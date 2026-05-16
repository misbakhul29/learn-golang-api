package main

import (
	"belajargo/internal/database"
	"belajargo/internal/services"
	"fmt"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		services.Log(services.Logger{
			Name:    "SEEDER",
			Message: fmt.Sprintf("Failed to connect database: %v", err),
		})
		return
	}

	err = database.Seeders(db)
	if err == nil {
		services.Log(services.Logger{
			Name:    "SEEDER",
			Message: "Successfully seeded database",
		})
		return
	}
}
