package handlers

import (
	"book_keeper/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type testModel struct {
	person         models.Person
	expectedResult int
}

func TestPostPerson(t *testing.T) {
	f := []testModel{
		testModel{person: models.Person{
			Name:  "Manasses",
			Email: "manasses@",
			Books: nil,
		},
			expectedResult: 400,
		},
		testModel{person: models.Person{
			Name:  "",
			Email: "manasses@",
			Books: nil,
		},
			expectedResult: 400,
		},
		testModel{person: models.Person{
			Name:  "Manasses julio",
			Email: "manasses@hotmail.com",
			Books: nil,
		},
			expectedResult: 200,
		},
	}

	for _, test := range f {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(test.person)
		if err != nil {
			log.Fatal(err)
		}
		r, _ := http.NewRequest("POST", "/person", &buf)
		w := httptest.NewRecorder()

		Routes().ServeHTTP(w, r)
		assert.Equal(t, test.expectedResult, w.Code)
	}

}

func TestPutPerson(t *testing.T) {
	f := []models.Person{
		{
			Name:  "Manasses",
			Email: "manasses@",
			Books: nil,
		},
		{
			Name:  "Manasses Júlio",
			Email: "manasses@.com",
			Books: nil,
		},
		{
			Name:  "",
			Email: "manasses@.com",
			Books: nil,
		},
	}

	for id, personUpdate := range f {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(personUpdate)
		if err != nil {
			log.Fatal(err)
		}

		url := "/person/" + strconv.Itoa(id+1)

		r, _ := http.NewRequest("PUT", url, &buf)
		w := httptest.NewRecorder()

		Routes().ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	}
}
