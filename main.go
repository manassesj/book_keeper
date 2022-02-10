package main

import (
	database "book_keeper/db"
	"book_keeper/models"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
	"net/http"
)

var (
	person = &models.Person{
		Name:  "Nome1",
		Email: "NOME1@hotmail.com",
	}
	books = []models.Book{
		{Title: "Book1", Author: "Author1", CallNumber: 1234, PersonID: 1},
		{Title: "Book2", Author: "Author2", CallNumber: 12345, PersonID: 1},
	}
)
var db *gorm.DB = database.GetDbManager().DB
var err error

func main() {
	//Loading enviroment variables
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.Book{})

	//Make migration to the database if they hae not alerady vbeen created
	db.Create(person)
	for idx := range books {
		db.Create(&books[idx])
	}

	// API routes

	router := mux.NewRouter()

	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/person/{id}", getPerson).Methods("GET")
	router.HandleFunc("/person", createPerson).Methods("POST")
	router.HandleFunc("/person/{id}", deletePerson).Methods("DELETE")
	router.HandleFunc("/person/{id}", updatePerson).Methods("UPDATE")

	http.ListenAndServe(":8080", router)
}

func getPeople(rw http.ResponseWriter, r *http.Request) {
	var people []models.Person
	db.First(&people)
	json.NewEncoder(rw).Encode(&people)
}

func getPerson(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person models.Person
	var books []models.Book
	db.First(&person, params["id"])
	db.Model(&person)
	person.Books = books
	json.NewEncoder(rw).Encode(person)
}

func createPerson(rw http.ResponseWriter, r *http.Request) {
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)
	createdPerson := db.Create(&person)
	err := createdPerson.Error
	if err != nil {
		json.NewEncoder(rw).Encode(err)
	} else {
		json.NewEncoder(rw).Encode(&person)
	}
}

func deletePerson(rw http.ResponseWriter, r *http.Request) {
	var person models.Person
	params := mux.Vars(r)

	db.First(&person, params["id"])
	db.Delete(&person)
	json.NewEncoder(rw).Encode(&person)
}

func updatePerson(rw http.ResponseWriter, r *http.Request) {
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)
	//db.Update(&person).Updates(&person)
	json.NewEncoder(rw).Encode(&person)
}
