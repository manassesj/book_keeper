package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name  string
	Email string `gorm:"type:varchar(100);unique_index"`
	Books []Book
}
