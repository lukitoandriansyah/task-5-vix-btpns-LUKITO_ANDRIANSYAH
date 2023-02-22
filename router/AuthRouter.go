package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/controllers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

var (
	db                  *gorm.DB
	usersRepo           models.UsersRepo
	authHelperInterface helpers.AuthHelperInterface
	jwtHelperInterface  helpers.JwtHelperInterface
	authInterface       controllers.AuthInterface
)

func AuthRouter() {
	router := gin.Default()
	authRoutes := router.Group("api/users")
	{
		authRoutes.POST("/login", authInterface.Login)
		authRoutes.POST("/register", authInterface.Register)
	}
}
