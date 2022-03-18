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
