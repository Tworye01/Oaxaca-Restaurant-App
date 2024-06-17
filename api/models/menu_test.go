package models

import (
	"team-project/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMenuAutoMigrate(t *testing.T) {
	s := database.New()
	err := s.AutoMigrate(&Menu{})
	assert.NoError(t, err)
}
