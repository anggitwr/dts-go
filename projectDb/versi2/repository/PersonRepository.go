package repository

import (
	"versi2/model"

	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{
		db: db,
	}
}

func (pr *PersonRepository) Create(newPerson model.Person) (model.Person, error) {

	tx := pr.db.Create(&newPerson)

	return newPerson, tx.Error
}

func (pr *PersonRepository) GetAll() ([]model.Person, error) {
	var persons = []model.Person{}

	tx := pr.db.Find(&persons)

	return persons, tx.Error
}
