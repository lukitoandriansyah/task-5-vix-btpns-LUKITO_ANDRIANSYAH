package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

/*func DBinit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:wj3oYiTnE1uySOd2r0$K^@tcp(127.0.0.1:3306)/photousers?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(models.Users{}, models.Photos{})
	return db
}
*/

func DataBaseConnection() *gorm.DB {
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
	db.AutoMigrate(&models.Users{}, &models.Photos{})
	return db
}
