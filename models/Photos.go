package models

import "gorm.io/gorm"

type Photos struct {
	//gorm.Model
	ID       uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Caption  string `gorm:"type:text" json:"caption"`
	PhotoUrl string `gorm:"type:text" json:"photoUrl"`
	UserId   uint   `gorm:"not null" json:"-"`
	Users    Users  `gorm:"foreignKey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}

type PhotosRepo interface {
	InserPhoto(photos Photos) Photos
	UpdatePhoto(photos Photos) Photos
	DeletePhoto(photos Photos) Photos
	GetPhotoById(photoID uint) Photos
}

type photoConnection struct {
	connection *gorm.DB
}

func (pc *photoConnection) InserPhoto(photos Photos) Photos {
	pc.connection.Save(&photos)
	pc.connection.Preload("Users").Find(&photos)
	return photos
}

func (pc *photoConnection) UpdatePhoto(photos Photos) Photos {
	pc.connection.Save(&photos)
	pc.connection.Preload("Users").Find(&photos)
	return photos
}

func (pc *photoConnection) DeletePhoto(photos Photos) Photos {
	pc.connection.Delete(&photos)
	pc.connection.Preload("Users").Find(&photos)
	return photos
}

func (pc *photoConnection) GetPhotoById(photoID uint) Photos {
	var photos Photos
	pc.connection.Preload("Users").Find(&photos, photoID)
	return photos
}

func NewPhotosRepo(db *gorm.DB) PhotosRepo {
	return &photoConnection{
		connection: db,
	}
}
