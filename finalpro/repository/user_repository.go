package repository

import (
	"finalpro/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data model.User) error
	Login(email string) (model.User, error)
	GetUserByID(id uint) (model.User, error)
	UpdateUser(data model.User) (model.User, error)
	DeleteUser(ID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(data model.User) error {
	err := r.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *userRepository) Login(email string) (model.User, error) {
	user := new(model.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}
func (r *userRepository) GetUserByID(id uint) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
func (r *userRepository) UpdateUser(data model.User) (model.User, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return model.User{}, err
	}
	return data, nil
}
func (r *userRepository) DeleteUser(ID uint) error {
	var user model.User
	user.ID = ID
	err := r.db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
