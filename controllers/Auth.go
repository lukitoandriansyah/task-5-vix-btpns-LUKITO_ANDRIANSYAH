package controllers

import "github.com/gin-gonic/gin"

type AuthInterface interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type AuthStruct struct {
}
