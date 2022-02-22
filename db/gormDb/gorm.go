package gormDb

import (
	"gorm.io/gorm"
)

type DbGorm struct {
	Db *gorm.DB
}

func (dbG *DbGorm) Create(i interface{}) error {
	e := dbG.Db.Create(i)
	return e.Error
}

func (dbG *DbGorm) Find(i interface{}) {
	dbG.Db.Find(i)
}

func (dbG *DbGorm) Update(i interface{}, id string) error {
	e := dbG.Db.Model(&i).Where("id = ?", id).Updates(i)
	return e.Error
}

func (dbG *DbGorm) Delete(i interface{}, id string) error {
	e := dbG.Db.Delete(&i, id)
	return e.Error
}
