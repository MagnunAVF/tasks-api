package db

import (
	"github.com/MagnunAVF/tasks-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error

	// Connect to the PostgreSQL database
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=mydatabase port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate DB
	if err := DB.AutoMigrate(&models.Task{}); err != nil {
		return err
	}

	return nil
}
