package model

import (
	"errors"
	"finalpro/helper"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Username  string `gorm:"not null;unique;type:varchar(100)" json:"username" form:"username" valid:"required~Your username is required"`
	Email     string `gorm:"not null;unique;type:varchar(100)" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required,minstringlength(6)~Password minimum lengths is 6 characters"`
	Age       uint   `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required"`
}

type ResponseUser struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Username  string `gorm:"not null;unique;type:varchar(100)" json:"username" form:"username" valid:"required~Your username is required"`
	Email     string `gorm:"not null;unique;type:varchar(100)" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required,minstringlength(6)~Password minimum lengths is 6 characters"`
	Age       uint   `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required"`
}

type RequestRegister struct {
	Username string `json:"username" example:"test"`
	Email    string `json:"email" example:"test@test.com"`
	Password string `json:"password" example:"123456"`
	Age      uint   `json:"age" example:"20"`
}

type RequestLogin struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"123456"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Age < 8 {
		err = errors.New("Minimum age to register is 8")
		return err
	}
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HashPass(u.Password)

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
