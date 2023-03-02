package router

import (
	"github.com/gin-gonic/gin"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/controllers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/middleware"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

func UsersRouter() {
	var db = database.Connection()
	var usersRepo = models.NewUsersRepo(db)
	var usersHelperInterface = helpers.NewUsersHelperInterface(usersRepo)
	var jwtHelperInterface = helpers.NewJwtHelperInterface()
	var usersInterface = controllers.NewUsersInterface(usersHelperInterface, jwtHelperInterface)
	router := gin.Default()

	usersRouter := router.Group("api/users", middleware.JwtAuth(jwtHelperInterface))
	{
		usersRouter.GET("/:userId", usersInterface.Profile)
		usersRouter.PUT("/:userId", usersInterface.Update)
	}
	router.Run()
}
