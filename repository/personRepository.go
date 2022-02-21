package repository

import (
	database "book_keeper/db"
	"book_keeper/models"
)

var db database.Db_interface = database.GetDbInstance()

func GetAllPeople(people []models.Person) []models.Person {
	db.Find(&people)
	return people
}

func AddPerson(person models.Person) error {
	err := db.Create(&person)
	return err
}

func UpdatePerson(person models.Person, id string) error {
	err := db.Update(person, id)
	return err
}
