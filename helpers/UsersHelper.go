package helpers

import (
	"github.com/mashingan/smapping"
	"log"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

type UsersHelperInterface interface {
	Update(users database.UsersUpdateData) models.Users
	Profile(userId string) models.Users
}

type UsersHelperStruct struct {
	usersRepo models.UsersRepo
}

func (uhs *UsersHelperStruct) Update(users database.UsersUpdateData) models.Users {
	userUpdate := models.Users{}
	err := smapping.FillStruct(&userUpdate, smapping.MapFields(&users))
	if err != nil {
		log.Fatalf("Failed to mapp #{err}")
	}
	updatedUser := uhs.usersRepo.UpdateUser(userUpdate)
	return updatedUser
}

func (uhs *UsersHelperStruct) Profile(userId string) models.Users {
	return uhs.usersRepo.ProfileUser(userId)
}

func NewUsersHelperInterface(usersHelperNew models.UsersRepo) UsersHelperInterface {
	return &UsersHelperStruct{
		usersRepo: usersHelperNew,
	}
}
