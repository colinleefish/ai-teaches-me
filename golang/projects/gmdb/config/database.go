package config

import (
	"fmt"
	"gmdb/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if GlobalConfig == nil {
		log.Fatal("Config not loaded. Call LoadConfig first.")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		GlobalConfig.Database.Host,
		GlobalConfig.Database.User,
		GlobalConfig.Database.Password,
		GlobalConfig.Database.DBName,
		GlobalConfig.Database.Port,
		GlobalConfig.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
	log.Printf("Connected to database: %s@%s:%d/%s",
		GlobalConfig.Database.User,
		GlobalConfig.Database.Host,
		GlobalConfig.Database.Port,
		GlobalConfig.Database.DBName,
	)

	// Run migrations
	RunMigrations()
}

// RunMigrations creates all tables from models
func RunMigrations() {
	log.Println("Running database migrations...")

	if err := DB.AutoMigrate(
		&models.Actor{},
		&models.Movie{},
		&models.Award{},
	); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Migrations completed successfully!")
}

func GetDB() *gorm.DB {
	return DB
}
