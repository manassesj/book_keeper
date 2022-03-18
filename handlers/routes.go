package handlers

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/person", CreatePerson).Methods("POST")
	router.HandleFunc("/person/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")
	router.HandleFunc("/person/{id}/books", GetAllBooks).Methods("GET")
	router.HandleFunc("/book", CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", DeleteBook).Methods("DELETE")
	return router
}
