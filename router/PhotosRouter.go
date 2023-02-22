package router

import (
	"github.com/gin-gonic/gin"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/controllers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/middleware"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

func PhotosRouter() {
	var photosRepo = models.NewPhotosRepo(database.Connection())
	var photosHelperInterface = helpers.NewPhotosHelperInterface(photosRepo)
	var jwtHelperInterface = helpers.NewJwtHelperInterface()
	var photosInterface = controllers.NewPhotosInterface(photosHelperInterface, jwtHelperInterface)

	router := gin.Default()

	photosRoute := router.Group("api/photos", middleware.JwtAuth(jwtHelperInterface))
	{
		photosRoute.POST("/photos", photosInterface.Insert)
		photosRoute.PUT("/:photoId", photosInterface.Update)
		photosRoute.GET("/photos", photosInterface.GetById)
		photosRoute.DELETE("/:photoId", photosInterface.Delete)
	}
	router.Run()
}
