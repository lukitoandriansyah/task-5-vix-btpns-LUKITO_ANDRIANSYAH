package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/helpers"
)

func JwtAuth(jwtHelperInterface helpers.JwtHelperInterface) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			res := helpers.BuildErrorResponse("Sorry, failed to process data", "Token not found", nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		token, err := jwtHelperInterface.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[userId]: ", claims["userId"])
			log.Println("Claim[issuer] ", claims["issuer"])
		} else {
			log.Println(err)
			res := helpers.BuildErrorResponse("Token is not valid", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, res)
		}
	}
}
