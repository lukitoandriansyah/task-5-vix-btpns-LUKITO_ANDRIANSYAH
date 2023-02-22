package main

import (
	"gorm.io/gorm"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/router"
)

var (
	db *gorm.DB = database.DataBaseConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)

	router.AuthRouter()
	router.PhotosRouter()
	router.UsersRouter()
}
