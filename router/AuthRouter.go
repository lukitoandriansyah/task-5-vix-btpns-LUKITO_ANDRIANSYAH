package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/controllers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

var (
	db                  *gorm.DB                    = database.DataBaseConnection()
	usersRepo           models.UsersRepo            = models.NewUsersRepo(db)
	authHelperInterface helpers.AuthHelperInterface = helpers.NewAuthHelper(usersRepo)
	jwtHelperInterface  helpers.JwtHelperInterface  = helpers.NewJwtHelperInterface()
	authInterface       controllers.AuthInterface   = controllers.NewAuthInterface(authHelperInterface, jwtHelperInterface)
)

func AuthRouter() {
	router := gin.Default()
	authRoutes := router.Group("api/users")
	{
		authRoutes.POST("/login", authInterface.Login)
		authRoutes.POST("/register", authInterface.Register)
	}
}
