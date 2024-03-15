package model

import (
	"gorm.io/gorm"
)

type Product struct {
	Price int
	gorm.Model
}
