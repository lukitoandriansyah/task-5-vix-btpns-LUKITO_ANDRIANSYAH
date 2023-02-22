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
	var login app.Login
	err := ctx.ShouldBind(&login)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to Process data", err.Error(), helpers.EmptyObjStruct{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authFix := as.authHelperInterface.VerifyCredential(login.Email, login.Password)
	if value, ok := authFix.(models.User); ok {
		generatedToken := as.jwtHelperInterface.GenerateToken(strconv.FormatUint(uint64(value.ID), 10))
		value.Token = generatedToken
		res := helpers.BuildResponse(true, "OK", value)
		ctx.JSON(http.StatusOK, res)
		return

	}
}

func (as AuthStruct) Register(ctx *gin.Context) {
	var register app.Register
	err := ctx.ShouldBind(&register)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed dto process data", err.Error(), helpers.EmptyObjStruct{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if !as.authHelperInterface.IsDuplicateEmail(register.Email) {
		res := helpers.BuildErrorResponse("failed to process data", err.Error(), helpers.EmptyObjStruct{})
		ctx.JSON(http.StatusConflict, res)
	} else {
		createUser := as.authHelperInterface.CreateUser(register)
		token := as.jwtHelperInterface.GenerateToken(strconv.FormatUint(uint64(createUser.ID), 10))
		createUser.Token = token
		res := helpers.BuildResponse(true, "OK!", createUser)
		ctx.JSON(http.StatusCreated, res)
	}
}

func NewAuthInterface(authHelperInterfaceNew helpers.AuthHelperInterface, jwtHelperInterfaceNew helpers.JwtHelperInterface) AuthInterface {
	return &AuthStruct{
		authHelperInterface: authHelperInterfaceNew,
		jwtHelperInterface:  jwtHelperInterfaceNew,
	}
}
