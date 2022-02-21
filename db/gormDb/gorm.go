package gormDb

import (
	"book_keeper/models"
	"gorm.io/gorm"
)

type DbGorm struct {
	Db *gorm.DB
}

func (dbG *DbGorm) Create(i interface{}) error {
	e := dbG.Db.Find(i)
	return e.Error
}

func (dbG *DbGorm) Find(i interface{}) {
	dbG.Db.Find(i)
}

func (dbG *DbGorm) Update(i interface{}, id string) error {
	e := dbG.Db.Model(&models.Person{}).Where("id = ?", id).Updates(i)
	return e.Error
}
