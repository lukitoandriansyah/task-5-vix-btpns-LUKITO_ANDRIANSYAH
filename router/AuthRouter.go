package router

import (
	"github.com/gin-gonic/gin"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/controllers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

/*
var (

	db                  = database.Connection()
	usersRepo           = models.NewUsersRepo(database.Connection())
	authHelperInterface = helpers.NewAuthHelper(usersRepo)
	jwtHelperInterface  = helpers.NewJwtHelperInterface()
	authInterface       = controllers.NewAuthInterface(authHelperInterface, jwtHelperInterface)

)
*/
func AuthRouter() {
	var db = database.Connection()
	var usersRepo = models.NewUsersRepo(db)
	var authHelperInterface = helpers.NewAuthHelper(usersRepo)
	var jwtHelperInterface = helpers.NewJwtHelperInterface()
	var authInterface = controllers.NewAuthInterface(authHelperInterface, jwtHelperInterface)
	router := gin.Default()
	authRoutes := router.Group("/api/users")
	{
		authRoutes.POST("/login", authInterface.Login)
		authRoutes.POST("/register", authInterface.Register)
	}
}
