package main

import (
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/router"
)

func main() {
	var db = database.Connection()
	defer database.CloseDatabaseConnection(db)

	router.AuthRouter()
	router.PhotosRouter()
	router.UsersRouter()
}
