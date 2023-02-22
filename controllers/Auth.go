package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/app"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

type AuthInterface interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type AuthStruct struct {
	authHelperInterface helpers.AuthHelperInterface
	jwtHelperInterface  helpers.JwtHelperInterface
}

func (as AuthStruct) Login(ctx *gin.Context) {
	var registerData app.Login
	err := ctx.ShouldBind(&registerData)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to Process data", err.Error(), helpers.EmptyObjStruct{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authFix := as.authHelperInterface.VerifyCredential(registerData.Email, registerData.Password)
	if value, ok := authFix.(models.Users); ok {
		generatedToken := as.jwtHelperInterface.GenerateToken(strconv.FormatUint(uint64(value.ID), 10))
		value.Token = generatedToken
		res := helpers.BuildResponse(true, "OK", value)
		ctx.JSON(http.StatusOK, res)
		return

	}
}

func (as AuthStruct) Register(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewAuthInterface(authHelperInterfaceNew helpers.AuthHelperInterface, jwtHelperInterfaceNew helpers.JwtHelperInterface) AuthInterface {
	return &AuthStruct{
		authHelperInterface: authHelperInterfaceNew,
		jwtHelperInterface:  jwtHelperInterfaceNew,
	}
}
