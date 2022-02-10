package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Dbmanager struct {
	DB *gorm.DB
}

var singleInstanceDb = new()

func GetDbManager() Dbmanager {
	return singleInstanceDb
}

func new() Dbmanager {
	host := os.Getenv("DBHOST")
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	password := os.Getenv("DBPASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s TimeZone=America/Sao_Paulo",
		host, user, dbName, password, dbport)

	localDb, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("Successfully connected to database")
	}

	return Dbmanager{DB: localDb}
}
