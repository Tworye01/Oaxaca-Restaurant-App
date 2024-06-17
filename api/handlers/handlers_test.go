package handlers

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	v := m.Run()
	os.Exit(v)
}
