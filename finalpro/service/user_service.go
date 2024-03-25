package service

import (
	"finalpro/helper"
	"finalpro/model"
	"finalpro/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(data *model.User) error
	Login(data *model.RequestLogin) (model.ResponseLogin, error)
	GetUserByID(ID uint) (model.ResponseUser, error)
	UpdateUser(data model.User, ID uint) (model.User, error)
	DeleteUser(ID uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(data *model.User) error {
	err := s.repo.Create(*data)
	if err != nil {
		return err
	}
	return nil
}
func (s *userService) Login(data *model.RequestLogin) (model.ResponseLogin, error) {
	dataUser, err := s.repo.Login(data.Email)
	if err != nil {
		return model.ResponseLogin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(data.Password))
	if err != nil {
		return model.ResponseLogin{}, err
	}

	token := helper.GenerateToken(dataUser.ID, dataUser.Email)
	if err != nil {
		return model.ResponseLogin{}, err
	}

	resp := model.ResponseLogin{}
	resp.Token = token

	return resp, nil
}

func (service *userService) GetUserByID(ID uint) (model.ResponseUser, error) {
	resUser, err := service.repo.GetUserByID(ID)
	if err != nil {
		return model.ResponseUser{}, err
	}
	var response model.ResponseUser
	response.ID = resUser.ID
	response.Username = resUser.Username
	response.Email = resUser.Email
	response.Age = resUser.Age

	return response, nil
}

func (service *userService) UpdateUser(data model.User, ID uint) (model.User, error) {
	entityUser := model.User{}
	entityUser.ID = uint(ID)
	entityUser.Email = data.Email
	entityUser.Username = data.Username
	entityUser.Age = data.Age
	entityUser.Password = data.Password
	update, err := service.repo.UpdateUser(entityUser)
	if err != nil {
		return model.User{}, err
	}
	return update, nil
}

func (service *userService) DeleteUser(ID uint) error {
	err := service.repo.DeleteUser(ID)
	if err != nil {
		return err
	}
	return nil
}
