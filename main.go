package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"os"
)

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int
	PersonID   int
}

var (
	person = &Person{
		Name:  "Nome1",
		Email: "NOME1@hotmail.com",
	}
	books = []Book{
		{Title: "Book1", Author: "Author1", CallNumber: 1234, PersonID: 1},
		{Title: "Book2", Author: "Author2", CallNumber: 12345, PersonID: 1},
	}
)
var db *gorm.DB
var err error

func main() {
	//Loading enviroment variables

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//Database conection string

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		host, user, dbName, password, dbport)

	//openning connection to database

	db, err = gorm.Open(dialect, dbURI)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	//CLose connection to database when the main function finises

	defer db.Close()

	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})

	//Make migration to the database if they hae not alerady vbeen created
	/*	db.Create(person)
		for idx := range books {
			db.Create(&books[idx])
		}
	*/

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
	var people []Person
	db.Find(&people)

	json.NewEncoder(rw).Encode(&people)
}

func getPerson(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person Person
	var books []Book
	db.First(&person, params["id"])
	db.Model(&person).Related(&books)
	person.Books = books
	json.NewEncoder(rw).Encode(person)
}

func createPerson(rw http.ResponseWriter, r *http.Request) {
	var person Person
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
	var person Person
	params := mux.Vars(r)

	db.First(&person, params["id"])
	db.Delete(&person)
	json.NewEncoder(rw).Encode(&person)
}

func updatePerson(rw http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	db.Update(&person).Updates(&person)
	json.NewEncoder(rw).Encode(&person)
}
