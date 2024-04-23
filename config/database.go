package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AliRamdhan/trainticket/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	dbUsername := os.Getenv("DB_USERNAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbDatabase := os.Getenv("DB_DATABASE")

	// Construct the DSN string using environment variables
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s", dbUsername, dbHost, dbPort, dbDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func AutoMigrate() error {
	return DB.AutoMigrate(&model.Role{}, &model.User{}, &model.Ticket{}, &model.Seat{}, &model.Order{}, &model.Payment{})
	// return DB.AutoMigrate(&model.Client{}, &model.Profile{})
}

// load seed data into the database
func SeedData() ([]model.User, []model.Role) {
	var roles = []model.Role{
		{Name: "admin", Description: "Administrator role"},
		{Name: "customer", Description: "Authenticated customer role"},
		{Name: "visitor", Description: "Unauthenticated customer role"},
	}
	var user = []model.User{
		{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleUser: 1, CreatedAt: time.Now().Format("2006-01-02 15:04:05")},
	}

	if err := DB.Save(&roles).Error; err != nil {
		log.Fatalf("Error saving roles: %v", err)
	}
	if err := DB.Save(&user).Error; err != nil {
		log.Fatalf("Error saving users: %v", err)
	}

	return user, roles
}
