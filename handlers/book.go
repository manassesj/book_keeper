package handlers

import (
	"book_keeper/models"
	"book_keeper/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllBooks(rw http.ResponseWriter, r *http.Request) {
	var books []models.Book

	books = repository.GetAllBooks(books)
	json.NewEncoder(rw).Encode(&books)
}

func CreateBook(rw http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	result := repository.AddBook(book)

	if result != nil {
		json.NewEncoder(rw).Encode(result)
	} else {
		json.NewEncoder(rw).Encode(&book)
	}
}

func UpdateBook(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	repository.UpdateBook(book, params["id"])
	json.NewEncoder(rw).Encode(&book)
}

func DeleteBook(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repository.DeleteBook(models.Book{}, params["id"])
}
