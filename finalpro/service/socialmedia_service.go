package service

import (
	"finalpro/model"
	"finalpro/repository"
)

type SocialMediaService interface {
	GetAllSosmeds() ([]model.ResponseSocialMedia, error)
	GetSosmedByID(socialMediaID uint) (model.ResponseSocialMedia, error)
	CreateSosmed(data model.SocialMedia) (model.SocialMedia, error)
	UpdateSosmed(data model.SocialMedia, socialMediaID uint) (model.SocialMedia, error)
	DeleteSosmed(socialMediaID uint) error
}

type socialMediaService struct {
	sosmedRepo repository.SocialMediaRepository
}

func NewSosmedService(sosmedRepo repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{sosmedRepo: sosmedRepo}
}

func (service *socialMediaService) GetAllSosmeds() ([]model.ResponseSocialMedia, error) {
	resSosmed, err := service.sosmedRepo.GetSosmed()

	if err != nil {
		return []model.ResponseSocialMedia{}, err
	}
	var response []model.ResponseSocialMedia
	for _, sosmed := range resSosmed {
		tempResp := model.ResponseSocialMedia{}
		tempResp.ID = sosmed.ID
		tempResp.Name = sosmed.Name
		tempResp.SocialMediaURL = sosmed.SocialMediaURL
		tempResp.User.Username = sosmed.User.Username
		tempResp.User.Email = sosmed.User.Email
		tempResp.CreatedAt = sosmed.CreatedAt
		tempResp.UpdatedAt = sosmed.UpdatedAt
		response = append(response, tempResp)
	}

	return response, nil
}

func (service *socialMediaService) GetSosmedByID(socialMediaID uint) (model.ResponseSocialMedia, error) {
	resPhotos, err := service.sosmedRepo.GetSosmedByID(socialMediaID)
	if err != nil {
		return model.ResponseSocialMedia{}, err
	}
	var response model.ResponseSocialMedia
	response.ID = resPhotos.ID
	response.Name = resPhotos.Name
	response.SocialMediaURL = resPhotos.SocialMediaURL
	response.User.Username = resPhotos.User.Username
	response.User.Email = resPhotos.User.Email
	response.CreatedAt = resPhotos.CreatedAt
	response.UpdatedAt = resPhotos.UpdatedAt
	return response, nil
}

func (service *socialMediaService) CreateSosmed(data model.SocialMedia) (model.SocialMedia, error) {
	create, err := service.sosmedRepo.CreateSosmed(data)
	if err != nil {
		return model.SocialMedia{}, err
	}
	return create, nil
}

func (service *socialMediaService) UpdateSosmed(data model.SocialMedia, socialMediaID uint) (model.SocialMedia, error) {
	entitySosmed := model.SocialMedia{}
	entitySosmed.ID = uint(socialMediaID)
	entitySosmed.Name = data.Name
	entitySosmed.SocialMediaURL = data.SocialMediaURL
	update, err := service.sosmedRepo.UpdateSosmed(entitySosmed)
	if err != nil {
		return model.SocialMedia{}, err
	}
	return update, nil
}

func (service *socialMediaService) DeleteSosmed(socialMediaID uint) error {
	err := service.sosmedRepo.DeleteSosmed(socialMediaID)
	if err != nil {
		return err
	}
	return nil
}
