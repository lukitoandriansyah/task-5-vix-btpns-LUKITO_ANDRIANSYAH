package controllers

import "task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"

/*
func (idb *InDB) GetUser(c *gin.Context) {
	var (
		users  models.Users
		result gin.H
	)
}*/

type PhotosInterface interface {
}

type PhotosStruct struct {
}

func NewPhotosInterface(photosHelperInterfaceNew helpers.PhotosHelperInterface, jwtHelperInterfaceNew helpers.JwtHelperInterface) PhotosInterface {
	return &PhotosStruct{}
}
