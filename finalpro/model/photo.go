package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         `gorm:"not null" json:"title" form:"title" valid:"required~Title is required" example:"A Photo"`
	Caption   string         `json:"caption" form:"caption" example:"My Photo"`
	PhotoURL  string         `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required" example:"https://example.com/photo.jpg"`
	UserID    uint
	User      *User `json:"omitempty"`
}
type RequestPhoto struct {
	Title    string `json:"title" example:"A Photo"`
	Caption  string `json:"caption,omitempty" example:"My Photo"`
	PhotoURL string `json:"photo_url" example:"https://example.com/photo.jpg"`
}
type ResponsePhoto struct {
	ID        uint   `json:"id,omitempty" example:"1"`
	Title     string `json:"title" example:"A Photo"`
	Caption   string `json:"caption,omitempty" example:"My Photo"`
	PhotoURL  string `json:"photo_url" example:"https://example.com/photo.jpg"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	User      struct {
		Username string `json:"username" example:"anon"`
		Email    string `json:"email" example:"test@example.com"`
	} `json:"user"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
