package helpers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JwtHelperInterface interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtHelperCustomStruct struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

type JwtHelperStruct struct {
	secretKey string
	issuer    string
}

func (jwtHelperStruct JwtHelperStruct) GenerateToken(userId string) string {
	claims := &JwtHelperCustomStruct{
		userId, jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    jwtHelperStruct.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenValue, err := token.SignedString([]byte(jwtHelperStruct.secretKey))
	if err != nil {
		panic(err)
	}
	return tokenValue
}

func (jwtHelperStruct JwtHelperStruct) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token_ *jwt.Token) (interface{}, error) {
		if _, ok := token_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpexted signing method %v", token_.Header["alg"])
		}
		return []byte(jwtHelperStruct.secretKey), nil
	})

}

func NewJwtHelperInterface() JwtHelperInterface {
	return &JwtHelperStruct{
		issuer:    "Lukito_21",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "lukito_21" {
		secretKey = "lukito_21"
	}
	return secretKey
}
