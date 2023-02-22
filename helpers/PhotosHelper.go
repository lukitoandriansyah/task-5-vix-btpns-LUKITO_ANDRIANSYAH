package helpers

import (
	"fmt"
	"github.com/mashingan/smapping"
	"log"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/database"
	"task-5-vix-btpns-LUKITO_ANDRIANSYAH/models"
)

type PhotosHelperInterface interface {
	Insert(photoDb database.PhotoCreateData) models.Photos
	Update(photoDb database.PhotosUpdateData) models.Photos
	Delete(photoDb models.Photos)
	GetById(photodId uint) models.Photos
	IsAllowedToEdit(usersId string, PhotosId uint) bool
}

type PhotosHelperStruct struct {
	photosRepo models.PhotosRepo
}

func (phs *PhotosHelperStruct) Insert(photoDb database.PhotoCreateData) models.Photos {
	photo := models.Photos{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&photoDb))
	if err != nil {
		log.Fatalf("Failed to map #{err}")
	}
	res := phs.photosRepo.InserPhoto(photo)
	return res
}

func (phs *PhotosHelperStruct) Update(photoDb database.PhotosUpdateData) models.Photos {
	photo := models.Photos{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&photoDb))
	if err != nil {
		log.Fatalf("Failed to  mapp #{err}")
	}
	res := phs.photosRepo.UpdatePhoto(photo)
	return res
}

func (phs *PhotosHelperStruct) Delete(photoDb models.Photos) {
	phs.photosRepo.DeletePhoto(photoDb)
}

func (phs *PhotosHelperStruct) GetById(photodId uint) models.Photos {
	return phs.photosRepo.GetPhotoById(photodId)
}

func (phs *PhotosHelperStruct) IsAllowedToEdit(usersId string, PhotosId uint) bool {
	pid := phs.photosRepo.GetPhotoById(PhotosId)
	id := fmt.Sprintf(string(pid.UserId))
	return usersId == id
}

func NewPhotosHelperInterface(photosRepo models.PhotosRepo) PhotosHelperInterface {
	return &PhotosHelperStruct{
		photosRepo: photosRepo,
	}
}
