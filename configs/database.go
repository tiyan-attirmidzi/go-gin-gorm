package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/tiyan-attirmidzi/go-rest-api/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbName)
	db, errDB := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if errDB != nil {
		panic("Failed to Connect Database")
	}

	// TODO: Add Model to Migrate
	db.AutoMigrate(
		&entities.User{},
		&entities.Book{},
	)

	return db

}

func DatabaseDisconnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to Rejection Database")
	}

	dbSQL.Close()

}
