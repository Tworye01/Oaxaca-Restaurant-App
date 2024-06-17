package models

import (
	"team-project/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrinkAutoMigrate(t *testing.T) {
	s := database.New()
	err := s.AutoMigrate(&Drink{})
	assert.NoError(t, err)
}
