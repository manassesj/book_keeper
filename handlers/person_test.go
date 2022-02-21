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

func TestPostPerson(t *testing.T) {
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

	for _, test := range f {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(test)
		if err != nil {
			log.Fatal(err)
		}
		r, _ := http.NewRequest("POST", "/person", &buf)
		w := httptest.NewRecorder()

		Routes().ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
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
