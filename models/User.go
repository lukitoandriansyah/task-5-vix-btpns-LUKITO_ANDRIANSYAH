package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID        uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string   `gorm:"type:varchar(255)" json:"username"`
	Email     string   `gorm:"uniqueIndex; type:varchar(255)" json:"email"`
	Password  string   `gorm:"->'<-; not null" validate:"min=6" json:"."`
	Token     string   `gorm:"-" json:"token,omitempty"`
	Photos    *[]Photo `json:"photos,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsersRepo interface {
	InsertUser(users User) User
	UpdateUser(users User) User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (ab *gorm.DB)
	FindByEmail(email string) User
	GetUser(users User) User
	DeleteUser(users User) User
	ProfileUser(userId string) User
}

type userConnection struct {
	connection *gorm.DB
}

func (uc userConnection) InsertUser(users User) User {
	users.Password = hashAndSalt([]byte(users.Password))
	uc.connection.Save(&users)
	return users
}

func (uc userConnection) UpdateUser(users User) User {
	if users.Password != "" {
		users.Password = hashAndSalt([]byte(users.Password))
	} else {
		var tempUsers User
		uc.connection.Find(&tempUsers, users.ID)
		users.Password = tempUsers.Password
	}
	uc.connection.Save(&users)
	return users
}

func (uc userConnection) VerifyCredential(email string, password string) interface{} {
	var users User
	res := uc.connection.Where("email=?", email).Take(&users)
	if res.Error != nil {
		return users
	}
	return nil
}

func (uc userConnection) IsDuplicateEmail(email string) (ab *gorm.DB) {
	var users User
	return uc.connection.Where("email = ?", email).Take(&users)
}

func (uc userConnection) FindByEmail(email string) User {
	var users User
	uc.connection.Where("email=?", email).Take(&users)
	return users
}

func (uc userConnection) GetUser(users User) User {
	//var users User
	uc.connection.Preload("User").Find(&users)
	return users
}

func (uc userConnection) DeleteUser(users User) User {
	uc.connection.Delete(&users)
	return users
}

func (uc userConnection) ProfileUser(userId string) User {
	var users User
	uc.connection.Preload("Photo").Preload("Photo.User").Find("&users, userId")
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
