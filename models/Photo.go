package models

import "gorm.io/gorm"

type Photo struct {
	//gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Caption  string `gorm:"type:text" json:"caption"`
	PhotoUrl string `gorm:"type:text" json:"photoUrl"`
	UserID   uint   `gorm:"not null" json:"-"`
	Users    User   `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}

type PhotosRepo interface {
	InsertPhoto(photos Photo) Photo
	UpdatePhoto(photos Photo) Photo
	DeletePhoto(photos Photo) Photo
	GetPhotoById(photoID uint) Photo
}

type photoConnection struct {
	connection *gorm.DB
}

func (pc *photoConnection) InsertPhoto(photos Photo) Photo {
	pc.connection.Save(&photos)
	pc.connection.Preload("User").Find(&photos)
	return photos
}

func (pc *photoConnection) UpdatePhoto(photos Photo) Photo {
	pc.connection.Save(&photos)
	pc.connection.Preload("User").Find(&photos)
	return photos
}

func (pc *photoConnection) DeletePhoto(photos Photo) Photo {
	pc.connection.Delete(&photos)
	pc.connection.Preload("User").Find(&photos)
	return photos
}

func (pc *photoConnection) GetPhotoById(photoID uint) Photo {
	var photos Photo
	pc.connection.Preload("User").Find(&photos, photoID)
	return photos
}

func NewPhotosRepo(db *gorm.DB) PhotosRepo {
	return &photoConnection{
		connection: db,
	}
}
