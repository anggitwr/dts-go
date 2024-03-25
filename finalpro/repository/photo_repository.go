package repository

import (
	"finalpro/model"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	GetPhotos() ([]model.Photo, error)
	GetPhotoByID(id uint) (model.Photo, error)
	CreatePhoto(data model.Photo) (model.Photo, error)
	UpdatePhoto(data model.Photo) (model.Photo, error)
	DeletePhoto(ID uint) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) GetPhotos() ([]model.Photo, error) {
	var photo []model.Photo
	err := r.db.Preload("User").Find(&photo).Error
	if err != nil {
		return []model.Photo{}, err
	}
	return photo, nil
}

func (r *photoRepository) GetPhotoByID(id uint) (model.Photo, error) {
	var photo model.Photo
	err := r.db.Preload("User").Where("id = ?", id).First(&photo).Error
	if err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}

func (r *photoRepository) CreatePhoto(data model.Photo) (model.Photo, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return model.Photo{}, err
	}
	return data, nil
}

func (r *photoRepository) UpdatePhoto(data model.Photo) (model.Photo, error) {
	err := r.db.Model(&data).Where("id = ?", data.ID).Updates(model.Photo{Title: data.Title, Caption: data.Caption, PhotoURL: data.PhotoURL}).Error
	if err != nil {
		return model.Photo{}, err
	}
	return data, nil
}

func (r *photoRepository) DeletePhoto(ID uint) error {
	photo := model.Photo{}
	photo.ID = uint(ID)
	err := r.db.Where("id = ?", ID).Delete(&photo).Error
	if err != nil {
		return err
	}
	return nil
}
