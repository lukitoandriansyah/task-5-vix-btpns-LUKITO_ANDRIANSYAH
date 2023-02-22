package controllers

import (
	"github.com/gin-gonic/gin"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
)

type UsersInterface interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type UsersStruct struct {
	usersHelperInterface helpers.UsersHelperInterface
	jwtHelperInterface   helpers.JwtHelperInterface
}

func (us *UsersStruct) Update(ctx *gin.Context) {

}

func (us *UsersStruct) Profile(ctx *gin.Context) {

}

func NewUsersInterface(usersHelperInterfaceNew helpers.UsersHelperInterface, jwtHelperInterfaceNew helpers.JwtHelperInterface) UsersInterface {
	return &UsersStruct{
		usersHelperInterface: usersHelperInterfaceNew,
		jwtHelperInterface:   jwtHelperInterfaceNew,
	}
}
