package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

func Connection() *gorm.DB {
	if godotenv.Load() != nil {
		panic("Failed to load env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dataConfig := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&local=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dataConfig), &gorm.Config{})

	if err != nil {
		panic("Failed to connecting database")
	}
	err = db.AutoMigrate(&models.User{}, &models.Photo{})
	return db
}
func CloseDatabaseConnection(db *gorm.DB) {
	dataSQL, err := db.DB()
	if err != nil {
		panic("Failed to disconnecting db")
	}
	err = dataSQL.Close()
}
