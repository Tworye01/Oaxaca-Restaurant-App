package models

import (
	"team-project/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCourse(t *testing.T) {
	str, err := Starter.String()
	assert.NoError(t, err)
	assert.Equal(t, "Starter", str)

	str, err = Main.String()
	assert.NoError(t, err)
	assert.Equal(t, "Main", str)

	str, err = Dessert.String()
	assert.NoError(t, err)
	assert.Equal(t, "Dessert", str)

	str, err = Side.String()
	assert.NoError(t, err)
	assert.Equal(t, "Side", str)

	str, err = Extra.String()
	assert.NoError(t, err)
	assert.Equal(t, "Extra", str)

	course := Course(100)
	_, err = course.String()
	assert.Error(t, err)
}

func TestSpice(t *testing.T) {
	str, err := Mild.String()
	assert.NoError(t, err)
	assert.Equal(t, "Mild", str)

	str, err = Medium.String()
	assert.NoError(t, err)
	assert.Equal(t, "Medium", str)

	str, err = Hot.String()
	assert.NoError(t, err)
	assert.Equal(t, "Hot", str)

	str, err = ExtraHot.String()
	assert.NoError(t, err)
	assert.Equal(t, "Extra Hot", str)

	spice := Spice(100)
	_, err = spice.String()
	assert.Error(t, err)
}

func TestDishAutoMigrate(t *testing.T) {
	s := database.New()
	err := s.AutoMigrate(&Dish{})
	assert.NoError(t, err)
}
