package helpers

import (
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/app"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

type AuthHelperInterface interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user app.Register) models.User
	FindByEmail(email string) models.User
	IsDuplicateEmail(email string) bool
}

type AuthHelperStruct struct {
	userRepo models.UsersRepo
}

func (authHelperStruct *AuthHelperStruct) VerifyCredential(email string, password string) interface{} {
	res := authHelperStruct.userRepo.VerifyCredential(email, password)
	if value, ok := res.(models.User); ok {
		comparedPassword := comparedPassword(value.Password, []byte(password))
		if value.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func comparedPassword(password string, bytes []byte) bool {
	bytesPassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(bytesPassword, bytes)
	if err != nil {
		log.Println(err)
		return false
	}
	return false
}

func (authHelperStruct AuthHelperStruct) CreateUser(user app.Register) models.User {
	userCreate := models.User{}
	err := smapping.FillStruct(&userCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed to map #{err}")
	}
	res := authHelperStruct.userRepo.InsertUser(userCreate)
	return res
}

func (authHelperStruct AuthHelperStruct) FindByEmail(email string) models.User {
	return authHelperStruct.userRepo.FindByEmail(email)
}

func (authHelperStruct AuthHelperStruct) IsDuplicateEmail(email string) bool {
	res := authHelperStruct.userRepo.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func NewAuthHelper(userApp models.UsersRepo) AuthHelperInterface {
	return &AuthHelperStruct{
		userRepo: userApp,
	}

}
