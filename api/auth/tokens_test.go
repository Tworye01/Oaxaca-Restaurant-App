package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testKey = []byte("test-key")

func createValidToken(t *testing.T) string {
	username := "username123"
	password := "password"
	token, err := CreateToken(username, password)
	assert.NotNil(t, token)
	assert.NoError(t, err)
	return token
}

func TestCreateToken(t *testing.T) {
	token := createValidToken(t)
	assert.NotEmpty(t, token)
}

func TestVerifyToken(t *testing.T) {
	token := createValidToken(t)
	err := VerifyToken(token)
	assert.NoError(t, err)
}

func TestInvalidToken(t *testing.T) {
	err := VerifyToken("invalidtoken")
	assert.Error(t, err)
}
