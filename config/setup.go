package config

import (
	"fmt"
	"os"

	"github.com/fabiansdp/golang_rest_api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Failed to load environment variables")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	errJoinTable := db.SetupJoinTable(&models.Shop{}, "Dorayakis", &models.ShopDorayaki{})

	if errJoinTable != nil {
		panic("Failed to create join table")
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Shop{}, &models.Dorayaki{})

	DB = db
}
