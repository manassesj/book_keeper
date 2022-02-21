package handlers

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/person", CreatePerson).Methods("POST")
	router.HandleFunc("/person/{id}", UpdatePerson).Methods("PUT")
	/*	router.HandleFunc("/person/{id}", getPerson).Methods("GET")
		router.HandleFunc("/person/{id}", deletePerson).Methods("DELETE")

	*/
	return router
}
