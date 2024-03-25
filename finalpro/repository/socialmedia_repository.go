package repository

import (
	"finalpro/model"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	GetSosmed() ([]model.SocialMedia, error)
	GetSosmedByID(id uint) (model.SocialMedia, error)
	CreateSosmed(data model.SocialMedia) (model.SocialMedia, error)
	UpdateSosmed(data model.SocialMedia) (model.SocialMedia, error)
	DeleteSosmed(socialMediaID uint) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSosmedRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db: db}
}

func (r *socialMediaRepository) GetSosmed() ([]model.SocialMedia, error) {
	var sosmed []model.SocialMedia
	err := r.db.Preload("User").Find(&sosmed).Error
	if err != nil {
		return []model.SocialMedia{}, err
	}
	return sosmed, nil
}

func (r *socialMediaRepository) GetSosmedByID(id uint) (model.SocialMedia, error) {
	var sosmed model.SocialMedia
	err := r.db.Preload("User").Where("id = ?", id).First(&sosmed).Error
	if err != nil {
		return model.SocialMedia{}, err
	}
	return sosmed, nil
}

func (r *socialMediaRepository) CreateSosmed(data model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return model.SocialMedia{}, err
	}
	return data, nil
}

func (r *socialMediaRepository) UpdateSosmed(data model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return model.SocialMedia{}, err
	}
	return data, nil
}

func (r *socialMediaRepository) DeleteSosmed(socialMediaID uint) error {
	sosmed := model.SocialMedia{}
	sosmed.ID = uint(socialMediaID)
	err := r.db.Where("id = ?", socialMediaID).Delete(&sosmed).Error
	if err != nil {
		return err
	}
	return nil
}
