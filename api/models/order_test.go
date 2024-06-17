package models

import (
	"team-project/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderStatusValue(t *testing.T) {
	str, err := Unknown.String()
	assert.NoError(t, err)
	assert.Equal(t, "Unknown", str)

	str, err = Cancelled.String()
	assert.NoError(t, err)
	assert.Equal(t, "Cancelled", str)

	str, err = Cooking.String()
	assert.NoError(t, err)
	assert.Equal(t, "Cooking", str)

	str, err = Serving.String()
	assert.NoError(t, err)
	assert.Equal(t, "Serving", str)

	Status := Status(100)
	_, err = Status.String()
	assert.Error(t, err)
}

func TestOrderAutoMigrate(t *testing.T) {
	s := database.New()
	err := s.AutoMigrate(&Order{})
	assert.NoError(t, err)
}
