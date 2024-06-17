package models

import (
	"team-project/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBillAutoMigrate(t *testing.T) {
	s := database.New()
	err := s.AutoMigrate(&Bill{})
	assert.NoError(t, err)
}
