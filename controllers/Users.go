package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
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
	var uud database.UsersUpdateData
	errData := ctx.ShouldBind(&uud)
	if errData != nil {
		res := helpers.BuildErrorResponse("Failed to process data", errData.Error(), helpers.EmptyObjStruct{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	tokenVal, errToken := us.jwtHelperInterface.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := tokenVal.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["userId"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	uud.ID = uint(id)
	updated := us.usersHelperInterface.Update(uud)
	res := helpers.BuildResponse(true, "OK!", updated)
	ctx.JSON(http.StatusOK, res)
}

func (us *UsersStruct) Profile(ctx *gin.Context) {

}

func NewUsersInterface(usersHelperInterfaceNew helpers.UsersHelperInterface, jwtHelperInterfaceNew helpers.JwtHelperInterface) UsersInterface {
	return &UsersStruct{
		usersHelperInterface: usersHelperInterfaceNew,
		jwtHelperInterface:   jwtHelperInterfaceNew,
	}
}
