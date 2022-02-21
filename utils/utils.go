package utils

import (
	"book_keeper/models"
	"strings"
)

func CheckPerson(person models.Person) bool {
	if CheckPersonName(person.Name) && CheckEmail(person.Email) {
		return true
	}
	return false
}

func CheckPersonName(name string) bool {
	if (len(name) >= 10 && len(name) <= 50) && countWords(name) >= 2 {
		return true
	}
	return false
}

func countWords(s string) int {
	return len(strings.Fields(s))
}

func CheckEmail(email string) bool {
	specCarac := "@"
	dotCom := ".com"

	if strings.Contains(email, specCarac) &&
		strings.Contains(email, dotCom) {
		return true
	}
	return false
}
