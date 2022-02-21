package utils

import "book_keeper/models"

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
