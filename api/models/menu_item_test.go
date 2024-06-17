package models

import (
	"team-project/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllergensScan(t *testing.T) {
	var a Allergens
	err := a.Scan("Wheat,Nuts")
	assert.NoError(t, err)
	assert.Equal(t, Allergens{"Wheat", "Nuts"}, a)

	var b Allergens
	err = b.Scan(123)
	assert.Error(t, err)
}

func TestAllergensValue(t *testing.T) {
	a := Allergens{"Wheat", "Nuts"}
	s, err := a.Value()
	assert.NoError(t, err)
	assert.Equal(t, "Wheat,Nuts", s)

	a = Allergens{}
	s, err = a.Value()
	assert.NoError(t, err)
	assert.Nil(t, s)
}

func TestMenuItemAutoMigrate(t *testing.T) {
	s := database.New()
	err := s.AutoMigrate(&MenuItem{})
	assert.NoError(t, err)
}
