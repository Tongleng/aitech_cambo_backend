package configs

import (
	"backend/models"
	"fmt"
	"log"
)

func RunMigrations() {
	fmt.Println("⏳ Checking database schema...")

	err := DB.AutoMigrate(
		&models.User{},
		&models.ProductCategory{},
		&models.SocialMedia{},
		&models.Product{},
	)

	if err != nil {
		log.Fatalf("Database Migration Failed: %v", err)
	}

	fmt.Println("Database Migrated Successfully")
}
