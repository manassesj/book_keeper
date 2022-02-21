package main

import (
	"book_keeper/handlers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

var err error

func main() {
	router := handlers.Routes()
	http.ListenAndServe(":8080", router)
}

/*func getPerson(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person models.Person
	var books []models.Book
	db.First(&person, params["id"])
	db.Model(&person)
	person.Books = books
	json.NewEncoder(rw).Encode(person)t
}*/

/*func deletePerson(rw http.ResponseWriter, r *http.Request) {
	var person models.Person
	params := mux.Vars(r)

	db.First(&person, params["id"])
	db.Delete(&person)
	json.NewEncoder(rw).Encode(&person)
}*/
