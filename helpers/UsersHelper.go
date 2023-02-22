package helpers

import (
	"github.com/mashingan/smapping"
	"log"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

type UsersHelperInterface interface {
	Update(users database.UsersUpdateData) models.User
	Profile(userId string) models.User
}

type UsersHelperStruct struct {
	usersRepo models.UsersRepo
}

func (uhs *UsersHelperStruct) Update(users database.UsersUpdateData) models.User {
	userUpdate := models.User{}
	err := smapping.FillStruct(&userUpdate, smapping.MapFields(&users))
	if err != nil {
		log.Fatalf("Failed to mapp #{err}")
	}
	updatedUser := uhs.usersRepo.UpdateUser(userUpdate)
	return updatedUser
}

func (uhs *UsersHelperStruct) Profile(userId string) models.User {
	return uhs.usersRepo.ProfileUser(userId)
}

func NewUsersHelperInterface(usersHelperNew models.UsersRepo) UsersHelperInterface {
	return &UsersHelperStruct{
		usersRepo: usersHelperNew,
	}
}
