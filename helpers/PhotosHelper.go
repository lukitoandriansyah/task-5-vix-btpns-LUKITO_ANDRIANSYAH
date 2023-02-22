package helpers

import (
	"fmt"
	"github.com/mashingan/smapping"
	"log"
	"strconv"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

type PhotosHelperInterface interface {
	Insert(photoDb database.PhotoCreateData) models.Photo
	Update(photoDb database.PhotosUpdateData) models.Photo
	Delete(photoDb models.Photo)
	GetById(photodId uint) models.Photo
	IsAllowedToEdit(usersId string, PhotosId uint) bool
}

type PhotosHelperStruct struct {
	photosRepo models.PhotosRepo
}

func (phs *PhotosHelperStruct) Insert(photoDb database.PhotoCreateData) models.Photo {
	photo := models.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&photoDb))
	if err != nil {
		log.Fatalf("Failed to map #{err}")
	}
	res := phs.photosRepo.InsertPhoto(photo)
	return res
}

func (phs *PhotosHelperStruct) Update(photoDb database.PhotosUpdateData) models.Photo {
	photo := models.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&photoDb))
	if err != nil {
		log.Fatalf("Failed to  mapp #{err}")
	}
	res := phs.photosRepo.UpdatePhoto(photo)
	return res
}

func (phs *PhotosHelperStruct) Delete(photoDb models.Photo) {
	phs.photosRepo.DeletePhoto(photoDb)
}

func (phs *PhotosHelperStruct) GetById(photodId uint) models.Photo {
	return phs.photosRepo.GetPhotoById(photodId)
}

func (phs *PhotosHelperStruct) IsAllowedToEdit(usersId string, PhotosId uint) bool {
	pid := phs.photosRepo.GetPhotoById(PhotosId)
	id := fmt.Sprintf(strconv.Itoa(int(pid.UserID)))
	return usersId == id
}

func NewPhotosHelperInterface(photosRepo models.PhotosRepo) PhotosHelperInterface {
	return &PhotosHelperStruct{
		photosRepo: photosRepo,
	}
}
