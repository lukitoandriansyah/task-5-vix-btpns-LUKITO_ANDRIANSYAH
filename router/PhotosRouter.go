package router

import (
	"github.com/gin-gonic/gin"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/controllers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/middleware"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

var (
	photosRepo            models.PhotosRepo             = models.NewPhotosRepo(db)
	photosHelperInterface helpers.PhotosHelperInterface = helpers.NewPhotosHelperInterface(photosRepo)
	/*jwtHelperInterface    helpers.JwtHelperInterface    = helpers.NewJwtHelperInterface()*/
	photosInterface controllers.PhotosInterface = controllers.NewPhotosInterface(photosHelperInterface, jwtHelperInterface)
)

func PhotosRouter() {
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
