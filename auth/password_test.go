package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordValidation(t *testing.T) {
	password := "password"
	dto, err := DigestPassword(password)
	assert.Nil(t, err, "Expected no err")
	assert.True(t, ValidatePassword(password, dto), "Expected valid password")
}
