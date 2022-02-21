package handlers

import (
	"book_keeper/models"
	"book_keeper/repository"
	"book_keeper/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPeople(rw http.ResponseWriter, r *http.Request) {
	var people []models.Person
	people = repository.GetAllPeople(people)
	json.NewEncoder(rw).Encode(&people)
}

func CreatePerson(rw http.ResponseWriter, r *http.Request) {
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)
	if utils.CheckPerson(person) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Success"))
		return
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Fail"))
		return
	}
	result := repository.AddPerson(person)

	if result != nil {
		json.NewEncoder(rw).Encode(result)
	} else {
		json.NewEncoder(rw).Encode(&person)
	}
}

func UpdatePerson(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)
	if utils.CheckPerson(person) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Success"))
		return
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Fail"))
		return
	}
	repository.UpdatePerson(person, params["id"])
	json.NewEncoder(rw).Encode(&person)
}
