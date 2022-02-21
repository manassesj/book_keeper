package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmail(t *testing.T) {

	emails := []string{
		"",
		"emailteste@",
		"email.com",
		"emailteste@.com",
		"email",
	}

	for _, email := range emails {
		value := CheckEmail(email)
		assert.Equal(t, true, value)
	}
}
