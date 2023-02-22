package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

type PhotosInterface interface {
	GetById(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type PhotosStruct struct {
	photosHelperInterface helpers.PhotosHelperInterface
	jwtHelperInterface    helpers.JwtHelperInterface
}

func (ps *PhotosStruct) GetById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("Parameter ID was not found", err.Error(), helpers.EmptyObjStruct{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var photos = ps.photosHelperInterface.GetById(uint(id))
	if (photos == models.Photo{}) {
		res := helpers.BuildErrorResponse("Data was not found", "There's not data for these id", helpers.EmptyObjStruct{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", photos)
		ctx.JSON(http.StatusOK, res)
	}
}

func (ps *PhotosStruct) Insert(ctx *gin.Context) {
	var pcd database.PhotoCreateData
	errData := ctx.ShouldBind(&pcd)
	if errData != nil {
		res := helpers.BuildErrorResponse("Failed when process data", errData.Error(), helpers.EmptyObjStruct{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userId := ps.getUserIdByToken(authHeader)
		convertUserId, err := strconv.ParseUint(userId, 10, 6)
		if err == nil {
			pcd.UserId = uint(convertUserId)
		}
		result := ps.photosHelperInterface.Insert(pcd)
		res := helpers.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusCreated, res)
	}
}

func (ps *PhotosStruct) Update(ctx *gin.Context) {
	var pud database.PhotosUpdateData
	errData := ctx.ShouldBind(&pud)
	if errData != nil {
		res := helpers.BuildErrorResponse("Failed to bproces data", errData.Error(), helpers.EmptyObjStruct{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	tokenVal, errToken := ps.jwtHelperInterface.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := tokenVal.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["userId"])
	if ps.photosHelperInterface.IsAllowedToEdit(userId, pud.ID) {
		id, errId := strconv.ParseUint(userId, 10, 10)
		if errId == nil {
			pud.UserId = uint(id)
		}
		result := ps.photosHelperInterface.Update(pud)
		res := helpers.BuildResponse(true, "OK!", result)
		ctx.JSON(http.StatusOK, res)
	} else {
		res := helpers.BuildErrorResponse("You don't have permission to acces this app", "you're not owner", helpers.EmptyObjStruct{})
		ctx.JSON(http.StatusForbidden, res)
	}
}

func (ps *PhotosStruct) Delete(ctx *gin.Context) {
	var photos models.Photo
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to get id", "parameter id was not exist", helpers.EmptyObjStruct{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	photos.ID = uint(id)
	authHeader := ctx.GetHeader("Authorization")
	tokenVal, errToken := ps.jwtHelperInterface.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := tokenVal.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["userId"])
	if ps.photosHelperInterface.IsAllowedToEdit(userId, photos.ID) {
		ps.photosHelperInterface.Delete(photos)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObjStruct{})
		ctx.JSON(http.StatusOK, res)
	} else {
		res := helpers.BuildErrorResponse("You don't have permission to access this", "you're not owner", helpers.EmptyObjStruct{})
		ctx.JSON(http.StatusOK, res)
	}
}

func NewPhotosInterface(photosHelperInterfaceNew helpers.PhotosHelperInterface, jwtHelperInterfaceNew helpers.JwtHelperInterface) PhotosInterface {
	return &PhotosStruct{
		photosHelperInterface: photosHelperInterfaceNew,
		jwtHelperInterface:    jwtHelperInterfaceNew,
	}
}

func (ps *PhotosStruct) getUserIdByToken(token string) string {
	valToken, err := ps.jwtHelperInterface.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := valToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["userId"])
	return id
}
