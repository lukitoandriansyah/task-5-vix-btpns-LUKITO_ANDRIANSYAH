package router

import (
	"github.com/jinzhu/gorm"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

var (
	db                  *gorm.DB
	usersRepo           models.UsersRepo
	authHelperInterface helpers.AuthHelperInterface
	jwtHelperInterface  helpers.JwtHelperInterface
)

func AuthRouter() {

}
