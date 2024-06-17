package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGileFileExtension(t *testing.T) {
	res, err := GetFileExtension("img.jpg")
	assert.NoError(t, err)
	assert.Equal(t, res, "jpg")

	_, err = GetFileExtension("img")
	assert.Error(t, err)
}
