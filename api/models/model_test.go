package models

import (
	"os"
	"team-project/database"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	v := m.Run()
	os.Exit(v)
}

func TestModelAutoMigrate(t *testing.T) {
	s := database.New()
	err := s.AutoMigrate(&Model{})
	assert.NoError(t, err)
}
