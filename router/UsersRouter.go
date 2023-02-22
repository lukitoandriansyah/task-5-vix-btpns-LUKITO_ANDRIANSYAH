package router

import (
	"github.com/gin-gonic/gin"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/controllers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/middleware"
)

var (
	usersHelperInterface helpers.UsersHelperInterface = helpers.NewUsersHelperInterface(usersRepo)
	usersInterface       controllers.UsersInterface   = controllers.NewUsersInterface(usersHelperInterface, jwtHelperInterface)
)

func UsersRouter() {
	router := gin.Default()

	usersRouter := router.Group("api/users", middleware.JwtAuth(jwtHelperInterface))
	{
		usersRouter.GET("/:userId", usersInterface.Profile)
		usersRouter.PUT("/:userId", usersInterface.Update)
	}
	router.Run()
}
