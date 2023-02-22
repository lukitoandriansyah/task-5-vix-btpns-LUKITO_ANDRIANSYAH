package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type Users struct {
	ID        uint      `gorm:"primaryKey:autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Email     string    `gorm:"uniqueIndex; type:varchar(255)" json:"email"`
	Password  string    `gorm:"->'<-; not null" validate:"min=6" json:"."`
	Token     string    `gorm:"-" json:"token,omitempty"`
	Photos    *[]Photos `gorm:"embedded" json:"photos,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsersRepo interface {
	InsertUser(users Users) Users
	UpdateUser(users Users) Users
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (ab *gorm.DB)
	FindByEmail(email string) Users
	GetUser(users Users) Users
	DeleteUser(users Users) Users
	ProfileUser(userId string) Users
}

type userConnection struct {
	connection *gorm.DB
}

func (uc userConnection) InsertUser(users Users) Users {
	users.Password = hashAndSalt([]byte(users.Password))
	uc.connection.Save(&users)
	return users
}

func (uc userConnection) UpdateUser(users Users) Users {
	if users.Password != "" {
		users.Password = hashAndSalt([]byte(users.Password))
	} else {
		var tempUsers Users
		uc.connection.Find(&tempUsers, users.ID)
		users.Password = tempUsers.Password
	}
	uc.connection.Save(&users)
	return users
}

func (uc userConnection) VerifyCredential(email string, password string) interface{} {
	var users Users
	res := uc.connection.Where("email=?", email).Take(&users)
	if res.Error == nil {
		return users
	}
	return nil
}

func (uc userConnection) IsDuplicateEmail(email string) (ab *gorm.DB) {
	var users Users
	return uc.connection.Where("email = ?", email).Take(&users)
}

func (uc userConnection) FindByEmail(email string) Users {
	var users Users
	uc.connection.Where("email=?", email).Take(&users)
	return users
}

func (uc userConnection) GetUser(users Users) Users {
	//var users Users
	uc.connection.Preload("Users").Find(&users)
	return users
}

func (uc userConnection) DeleteUser(users Users) Users {
	uc.connection.Delete(&users)
	return users
}

func (uc userConnection) ProfileUser(userId string) Users {
	var users Users
	uc.connection.Preload("Photos").Preload("Photos.User").Find("&users, userId")
	return users
}

func NewUsersRepo(db *gorm.DB) UsersRepo {
	return &userConnection{
		connection: db,
	}
}

func hashAndSalt(pass []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Fail to Hash Password")
	}
	return string(hash)
}
