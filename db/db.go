package db

import (
	"book_keeper/db/gormDb"
	"book_keeper/db/mock"
	"book_keeper/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Dbmanager struct {
	DB Db_interface
}

type Db_interface interface {
	Find(interface{})
	Create(interface{}) error
	Update(interface{}, string) error
}

func NewMock() Db_interface {
	return &mock.DbMock{}
}

func NewGorm() Db_interface {
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

	localDb.AutoMigrate(&models.Person{})
	localDb.AutoMigrate(&models.Book{})

	/*	db.Create(person)
		for idx := range books {
			db.Create(&books[idx])
		}*/

	return &gormDb.DbGorm{Db: localDb}
}

var db_int Db_interface

func GetDbInstance() Db_interface {
	if db_int == nil {
		if os.Getenv("DBIMPL") == "mock" {
			db_int = NewMock()
		} else {
			db_int = NewGorm()
		}
	}
	return db_int
}
