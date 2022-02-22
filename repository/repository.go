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

func DeletePerson(person models.Person, id string) error {
	err := db.Delete(person, id)
	return err
}

func GetAllBooks(books []models.Book) []models.Book {
	db.Find(&books)
	return books
}

func AddBook(book models.Book) error {
	err := db.Create(&book)
	return err
}

func UpdateBook(book models.Book, id string) error {
	err := db.Update(book, id)
	return err
}

func DeleteBook(book models.Book, id string) error {
	err := db.Delete(book, id)
	return err
}
